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
	time "time"

	machinev1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	versioned "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gardener/machine-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/client/listers/machine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineDeploymentInformer provides access to a shared informer and lister for
// MachineDeployments.
type MachineDeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineDeploymentLister
}

type machineDeploymentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMachineDeploymentInformer constructs a new informer for MachineDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineDeploymentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMachineDeploymentInformer constructs a new informer for MachineDeployment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineDeploymentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().MachineDeployments(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineV1alpha1().MachineDeployments(namespace).Watch(options)
			},
		},
		&machinev1alpha1.MachineDeployment{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineDeploymentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineDeploymentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineDeploymentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machinev1alpha1.MachineDeployment{}, f.defaultInformer)
}

func (f *machineDeploymentInformer) Lister() v1alpha1.MachineDeploymentLister {
	return v1alpha1.NewMachineDeploymentLister(f.Informer().GetIndexer())
}
