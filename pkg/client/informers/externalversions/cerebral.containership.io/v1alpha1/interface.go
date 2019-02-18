/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/containership/cerebral/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AutoscalingEngines returns a AutoscalingEngineInformer.
	AutoscalingEngines() AutoscalingEngineInformer
	// AutoscalingGroups returns a AutoscalingGroupInformer.
	AutoscalingGroups() AutoscalingGroupInformer
	// AutoscalingPolicies returns a AutoscalingPolicyInformer.
	AutoscalingPolicies() AutoscalingPolicyInformer
	// MetricsBackends returns a MetricsBackendInformer.
	MetricsBackends() MetricsBackendInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// AutoscalingEngines returns a AutoscalingEngineInformer.
func (v *version) AutoscalingEngines() AutoscalingEngineInformer {
	return &autoscalingEngineInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// AutoscalingGroups returns a AutoscalingGroupInformer.
func (v *version) AutoscalingGroups() AutoscalingGroupInformer {
	return &autoscalingGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// AutoscalingPolicies returns a AutoscalingPolicyInformer.
func (v *version) AutoscalingPolicies() AutoscalingPolicyInformer {
	return &autoscalingPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// MetricsBackends returns a MetricsBackendInformer.
func (v *version) MetricsBackends() MetricsBackendInformer {
	return &metricsBackendInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
