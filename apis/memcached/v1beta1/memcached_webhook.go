/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//
// Generated by:
//
// operator-sdk create webhook --group memcached --version v1beta1 --kind Memcached --programmatic-validation --defaulting
//

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// MemcachedDefaults -
type MemcachedDefaults struct {
	ContainerImageURL string
}

var memcachedDefaults MemcachedDefaults

// log is for logging in this package.
var memcachedlog = logf.Log.WithName("memcached-resource")

// SetupMemcachedDefaults - initialize Memcached spec defaults for use with either internal or external webhooks
func SetupMemcachedDefaults(defaults MemcachedDefaults) {
	memcachedDefaults = defaults
	memcachedlog.Info("Memcached defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *Memcached) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-memcached-openstack-org-v1beta1-memcached,mutating=true,failurePolicy=fail,sideEffects=None,groups=memcached.openstack.org,resources=memcacheds,verbs=create;update,versions=v1beta1,name=mmemcached.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Memcached{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Memcached) Default() {
	memcachedlog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this Memcached spec
func (spec *MemcachedSpec) Default() {
	if spec.ContainerImage == "" {
		spec.ContainerImage = memcachedDefaults.ContainerImageURL
	}
	spec.MemcachedSpecCore.Default()
}

// Default - common validations go here (for the OpenStackControlplane which uses this one)
func (spec *MemcachedSpecCore) Default() {
	// nothing here
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-memcached-openstack-org-v1beta1-memcached,mutating=false,failurePolicy=fail,sideEffects=None,groups=memcached.openstack.org,resources=memcacheds,verbs=create;update,versions=v1beta1,name=vmemcached.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Memcached{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateCreate() (admission.Warnings, error) {
	memcachedlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateUpdate(_ runtime.Object) (admission.Warnings, error) {
	memcachedlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Memcached) ValidateDelete() (admission.Warnings, error) {
	memcachedlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
