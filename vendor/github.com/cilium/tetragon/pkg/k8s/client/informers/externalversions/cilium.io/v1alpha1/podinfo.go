// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	ciliumiov1alpha1 "github.com/cilium/tetragon/pkg/k8s/apis/cilium.io/v1alpha1"
	versioned "github.com/cilium/tetragon/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/cilium/tetragon/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/cilium/tetragon/pkg/k8s/client/listers/cilium.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodInfoInformer provides access to a shared informer and lister for
// PodInfo.
type PodInfoInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodInfoLister
}

type podInfoInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodInfoInformer constructs a new informer for PodInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodInfoInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodInfoInformer constructs a new informer for PodInfo type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodInfoInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV1alpha1().PodInfo(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV1alpha1().PodInfo(namespace).Watch(context.TODO(), options)
			},
		},
		&ciliumiov1alpha1.PodInfo{},
		resyncPeriod,
		indexers,
	)
}

func (f *podInfoInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodInfoInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podInfoInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov1alpha1.PodInfo{}, f.defaultInformer)
}

func (f *podInfoInformer) Lister() v1alpha1.PodInfoLister {
	return v1alpha1.NewPodInfoLister(f.Informer().GetIndexer())
}