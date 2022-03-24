/*
Copyright 2022 The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	versioned "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned"
	internalinterfaces "github.com/kubernetes-csi/external-snapshotter/client/v6/informers/externalversions/internalinterfaces"
	v1 "github.com/kubernetes-csi/external-snapshotter/client/v6/listers/volumesnapshot/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeSnapshotContentInformer provides access to a shared informer and lister for
// VolumeSnapshotContents.
type VolumeSnapshotContentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VolumeSnapshotContentLister
}

type volumeSnapshotContentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVolumeSnapshotContentInformer constructs a new informer for VolumeSnapshotContent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeSnapshotContentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotContentInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeSnapshotContentInformer constructs a new informer for VolumeSnapshotContent type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeSnapshotContentInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SnapshotV1().VolumeSnapshotContents().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SnapshotV1().VolumeSnapshotContents().Watch(context.TODO(), options)
			},
		},
		&volumesnapshotv1.VolumeSnapshotContent{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeSnapshotContentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeSnapshotContentInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeSnapshotContentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumesnapshotv1.VolumeSnapshotContent{}, f.defaultInformer)
}

func (f *volumeSnapshotContentInformer) Lister() v1.VolumeSnapshotContentLister {
	return v1.NewVolumeSnapshotContentLister(f.Informer().GetIndexer())
}
