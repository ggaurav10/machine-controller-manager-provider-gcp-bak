/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// AWSMachineClasses returns a AWSMachineClassInformer.
	AWSMachineClasses() AWSMachineClassInformer
	// AlicloudMachineClasses returns a AlicloudMachineClassInformer.
	AlicloudMachineClasses() AlicloudMachineClassInformer
	// AzureMachineClasses returns a AzureMachineClassInformer.
	AzureMachineClasses() AzureMachineClassInformer
	// GCPMachineClasses returns a GCPMachineClassInformer.
	GCPMachineClasses() GCPMachineClassInformer
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// MachineClasses returns a MachineClassInformer.
	MachineClasses() MachineClassInformer
	// MachineDeployments returns a MachineDeploymentInformer.
	MachineDeployments() MachineDeploymentInformer
	// MachineSets returns a MachineSetInformer.
	MachineSets() MachineSetInformer
	// MachineTemplates returns a MachineTemplateInformer.
	MachineTemplates() MachineTemplateInformer
	// OpenStackMachineClasses returns a OpenStackMachineClassInformer.
	OpenStackMachineClasses() OpenStackMachineClassInformer
	// PacketMachineClasses returns a PacketMachineClassInformer.
	PacketMachineClasses() PacketMachineClassInformer
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

// AWSMachineClasses returns a AWSMachineClassInformer.
func (v *version) AWSMachineClasses() AWSMachineClassInformer {
	return &aWSMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AlicloudMachineClasses returns a AlicloudMachineClassInformer.
func (v *version) AlicloudMachineClasses() AlicloudMachineClassInformer {
	return &alicloudMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AzureMachineClasses returns a AzureMachineClassInformer.
func (v *version) AzureMachineClasses() AzureMachineClassInformer {
	return &azureMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPMachineClasses returns a GCPMachineClassInformer.
func (v *version) GCPMachineClasses() GCPMachineClassInformer {
	return &gCPMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineClasses returns a MachineClassInformer.
func (v *version) MachineClasses() MachineClassInformer {
	return &machineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineDeployments returns a MachineDeploymentInformer.
func (v *version) MachineDeployments() MachineDeploymentInformer {
	return &machineDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineSets returns a MachineSetInformer.
func (v *version) MachineSets() MachineSetInformer {
	return &machineSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MachineTemplates returns a MachineTemplateInformer.
func (v *version) MachineTemplates() MachineTemplateInformer {
	return &machineTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenStackMachineClasses returns a OpenStackMachineClassInformer.
func (v *version) OpenStackMachineClasses() OpenStackMachineClassInformer {
	return &openStackMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PacketMachineClasses returns a PacketMachineClassInformer.
func (v *version) PacketMachineClasses() PacketMachineClassInformer {
	return &packetMachineClassInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
