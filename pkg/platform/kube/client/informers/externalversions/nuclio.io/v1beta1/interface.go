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

package v1beta1

import (
	internalinterfaces "github.com/nuclio/nuclio/pkg/platform/kube/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// NuclioAPIGateways returns a NuclioAPIGatewayInformer.
	NuclioAPIGateways() NuclioAPIGatewayInformer
	// NuclioFunctions returns a NuclioFunctionInformer.
	NuclioFunctions() NuclioFunctionInformer
	// NuclioFunctionEvents returns a NuclioFunctionEventInformer.
	NuclioFunctionEvents() NuclioFunctionEventInformer
	// NuclioProjects returns a NuclioProjectInformer.
	NuclioProjects() NuclioProjectInformer
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

// NuclioAPIGateways returns a NuclioAPIGatewayInformer.
func (v *version) NuclioAPIGateways() NuclioAPIGatewayInformer {
	return &nuclioAPIGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NuclioFunctions returns a NuclioFunctionInformer.
func (v *version) NuclioFunctions() NuclioFunctionInformer {
	return &nuclioFunctionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NuclioFunctionEvents returns a NuclioFunctionEventInformer.
func (v *version) NuclioFunctionEvents() NuclioFunctionEventInformer {
	return &nuclioFunctionEventInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NuclioProjects returns a NuclioProjectInformer.
func (v *version) NuclioProjects() NuclioProjectInformer {
	return &nuclioProjectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
