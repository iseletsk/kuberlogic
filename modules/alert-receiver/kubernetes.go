/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	kuberlogicv1 "github.com/kuberlogic/kuberlogic/modules/operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	k8scheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"log"
)

var kuberLogicAlertCR = "kuberlogicalerts"
var kubeRestClient *rest.RESTClient

func initKubernetesClient() {
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Error initializing Kubernetes client: %s", err)
	}

	err = kuberlogicv1.AddToScheme(k8scheme.Scheme)
	if err != nil {
		log.Fatalf("Error adding clientset types to schema! %s", err)
	}

	crdConfig := *config
	crdConfig.ContentConfig.GroupVersion = &kuberlogicv1.GroupVersion
	crdConfig.APIPath = "/apis"
	crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(k8scheme.Scheme)
	crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()

	kubeRestClient, err = rest.UnversionedRESTClientFor(&crdConfig)
	if err != nil {
		log.Fatalf("Error initializing Kubernetes client: %s", err)
	}
}

func createAlertCR(name, namespace, alertName, alertValue, cluster, pod string) error {
	klAlert := kuberlogicv1.KuberLogicAlert{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: kuberlogicv1.KuberLogicAlertSpec{
			AlertName:  alertName,
			AlertValue: alertValue,
			Cluster:    cluster,
			Pod:        pod,
		},
	}

	res := kubeRestClient.Post().
		Namespace(namespace).
		Resource(kuberLogicAlertCR).
		Body(&klAlert).
		Do(context.TODO())
	return res.Error()
}

func deleteAlertCR(name, namespace string) error {
	res := kubeRestClient.Delete().
		Name(name).
		Namespace(namespace).
		Resource(kuberLogicAlertCR).
		Do(context.TODO())
	return res.Error()
}
