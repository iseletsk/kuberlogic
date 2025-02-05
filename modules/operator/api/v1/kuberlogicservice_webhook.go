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

package v1

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var log = logf.Log.WithName("kuberlogic-service-resource")

func (kls *KuberLogicService) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(kls).
		Complete()
}

var _ webhook.Defaulter = &KuberLogicService{}

func (kls *KuberLogicService) Default() {
	// TODO: figure out how to avoid cycle import and using versions into service-operator package
	var version string
	if kls.Spec.Type == "postgresql" {
		version = "13"
	} else if kls.Spec.Type == "mysql" {
		version = "5.7.26"
	}

	kls.InitDefaults(Defaults{
		VolumeSize: DefaultVolumeSize,
		Resources:  DefaultResources,
		Version:    version,
	})
}

var _ webhook.Validator = &KuberLogicService{}

type checkQuantity struct {
	current     resource.Quantity
	min         resource.Quantity
	errTemplate string
}

func (c *checkQuantity) compareQuanitity() error {
	if c.current.Cmp(c.min) < 0 {
		err := errors.New(fmt.Sprintf(c.errTemplate, c.min.String()))
		log.Error(err, "current less than min", "current", c.current.String(), "min", c.min)
		return err
	}
	return nil
}

// +kubebuilder:object:generate=false
type ErrorCollector []error

func (c *ErrorCollector) Collect(e error) {
	*c = append(*c, e)
}

func (c ErrorCollector) Error() (err string) {
	err = "Collected errors:\n"
	for _, e := range c {
		err += fmt.Sprintf("%s\n", e.Error())
	}
	return err
}

func (kls *KuberLogicService) ValidateCreate() error {
	errs := ErrorCollector{}
	checks := []checkQuantity{
		{
			current:     *kls.Spec.Resources.Limits.Cpu(),
			min:         *DefaultResources.Limits.Cpu(),
			errTemplate: "limits cpu is too small, min %s",
		},
		{
			current:     *kls.Spec.Resources.Limits.Memory(),
			min:         *DefaultResources.Limits.Memory(),
			errTemplate: "limits memory is too small, min %s",
		},
		{
			current:     *kls.Spec.Resources.Requests.Cpu(),
			min:         *DefaultResources.Requests.Cpu(),
			errTemplate: "requests cpu is too small, min %s",
		},
		{
			current:     *kls.Spec.Resources.Requests.Memory(),
			min:         *DefaultResources.Requests.Memory(),
			errTemplate: "requests memory is too small, min %s",
		},
		{
			current:     resource.MustParse(kls.Spec.VolumeSize),
			min:         resource.MustParse(DefaultVolumeSize),
			errTemplate: "volume size is too small, min %s",
		},
	}
	for _, c := range checks {
		if err := c.compareQuanitity(); err != nil {
			errs.Collect(err)
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (kls *KuberLogicService) ValidateUpdate(old runtime.Object) error {
	errs := ErrorCollector{}
	oldKls := old.(*KuberLogicService)
	if kls.Spec.Type != oldKls.Spec.Type {
		err := errors.New("type can not be changed")
		log.Error(err, "type can not be changed", "current", oldKls.Spec.Type, "new", kls.Spec.Type)
		errs.Collect(err)
	}

	currentVolume := resource.MustParse(oldKls.Spec.VolumeSize)
	newVolume := resource.MustParse(kls.Spec.VolumeSize)
	if newVolume.Cmp(currentVolume) < 0 {
		err := errors.New("volume size can not be decreased")
		log.Error(err, "volume size can not be decreased", "current", oldKls.Spec.VolumeSize, "new", kls.Spec.VolumeSize)
		errs.Collect(err)
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func (kls *KuberLogicService) ValidateDelete() error {
	log.Info("validate delete", "name", kls.Name)

	return nil
}
