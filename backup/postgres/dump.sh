#! /usr/bin/env bash

# enable unofficial bash strict mode
set -o errexit
set -o nounset
set -o pipefail
IFS=$'\n\t'

# integrate sentry
[[ -v SENTRY_DSN ]] && eval "$(sentry-cli bash-hook)"

ALL_DB_SIZE_QUERY="select sum(pg_database_size(datname)::numeric) from pg_database;"
PG_BIN=$PG_DIR/$PG_VERSION/bin
DUMP_SIZE_COEFF=5
ERRORCOUNT=0

BACKUP_NAME=$(date +%s).sql.gz
[[ ! -z "$DATABASE" ]] && BACKUP_NAME=$DATABASE-$BACKUP_NAME

TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
K8S_API_URL=https://$KUBERNETES_SERVICE_HOST:$KUBERNETES_SERVICE_PORT/api/v1
CERT=/var/run/secrets/kubernetes.io/serviceaccount/ca.crt

function estimate_size() {
  "$PG_BIN"/psql -tqAc "${ALL_DB_SIZE_QUERY}"
}

function dump() {
  if [ -n "$DATABASE" ]; then
    "$PG_BIN"/pg_dump $DATABASE
  else
    "$PG_BIN"/pg_dumpall
  fi
}

function compress() {
  pigz
}

function aws_upload() {
  declare -r EXPECTED_SIZE="$1"

  # mimic bucket setup from Spilo
  # to keep logical backups at the same path as WAL
  # NB: $LOGICAL_BACKUP_S3_BUCKET_SCOPE_SUFFIX already contains the leading "/" when set by the Postgres Operator
  PATH_TO_BACKUP=s3://$LOGICAL_BACKUP_S3_BUCKET"/postgresql/"$SCOPE$LOGICAL_BACKUP_S3_BUCKET_SCOPE_SUFFIX"/logical_backups/"$BACKUP_NAME

  args=()

  [[ ! -z "$EXPECTED_SIZE" ]] && args+=("--expected-size=$EXPECTED_SIZE")
  [[ ! -z "$LOGICAL_BACKUP_S3_ENDPOINT" ]] && args+=("--endpoint-url=$LOGICAL_BACKUP_S3_ENDPOINT")
  [[ ! -z "$LOGICAL_BACKUP_S3_REGION" ]] && args+=("--region=$LOGICAL_BACKUP_S3_REGION")
  [[ ! -z "$LOGICAL_BACKUP_S3_SSE" ]] && args+=("--sse=$LOGICAL_BACKUP_S3_SSE")

  aws s3 cp - "$PATH_TO_BACKUP" "${args[@]//\'/}"
}

function get_pods() {
  declare -r SELECTOR="$1"

  curl "${K8S_API_URL}/namespaces/${POD_NAMESPACE}/pods?$SELECTOR" \
    --cacert $CERT \
    -H "Authorization: Bearer ${TOKEN}" | jq .items[].status.podIP -r
}

function get_current_pod() {
  curl "${K8S_API_URL}/namespaces/${POD_NAMESPACE}/pods?fieldSelector=metadata.name%3D${HOSTNAME}" \
    --cacert $CERT \
    -H "Authorization: Bearer ${TOKEN}"
}

declare -a search_strategy=(
  list_all_replica_pods_current_node
  list_all_replica_pods_any_node
  get_master_pod
)

function list_all_replica_pods_current_node() {
  get_pods "labelSelector=${CLUSTER_NAME_LABEL}%3D${SCOPE},spilo-role%3Dreplica&fieldSelector=spec.nodeName%3D${CURRENT_NODENAME}" | head -n 1
}

function list_all_replica_pods_any_node() {
  get_pods "labelSelector=${CLUSTER_NAME_LABEL}%3D${SCOPE},spilo-role%3Dreplica" | head -n 1
}

function get_master_pod() {
  get_pods "labelSelector=${CLUSTER_NAME_LABEL}%3D${SCOPE},spilo-role%3Dmaster" | head -n 1
}

CURRENT_NODENAME=$(get_current_pod | jq .items[].spec.nodeName --raw-output)
export CURRENT_NODENAME

for search in "${search_strategy[@]}"; do

  PGHOST=$(eval "$search")
  export PGHOST

  if [ -n "$PGHOST" ]; then
    break
  fi

done

set -x
dump | compress | aws_upload $(($(estimate_size) / DUMP_SIZE_COEFF))
[[ ${PIPESTATUS[0]} != 0 || ${PIPESTATUS[1]} != 0 || ${PIPESTATUS[2]} != 0 ]] && ((ERRORCOUNT += 1))
set +x

exit $ERRORCOUNT
