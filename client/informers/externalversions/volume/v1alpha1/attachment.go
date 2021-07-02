/*
Copyright AppsCode Inc. and Contributors

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
	"context"
	time "time"

	volumev1alpha1 "kubeform.dev/provider-digitalocean-api/apis/volume/v1alpha1"
	versioned "kubeform.dev/provider-digitalocean-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-digitalocean-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-digitalocean-api/client/listers/volume/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AttachmentInformer provides access to a shared informer and lister for
// Attachments.
type AttachmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AttachmentLister
}

type attachmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAttachmentInformer constructs a new informer for Attachment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAttachmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAttachmentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAttachmentInformer constructs a new informer for Attachment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAttachmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumeV1alpha1().Attachments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumeV1alpha1().Attachments(namespace).Watch(context.TODO(), options)
			},
		},
		&volumev1alpha1.Attachment{},
		resyncPeriod,
		indexers,
	)
}

func (f *attachmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAttachmentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *attachmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumev1alpha1.Attachment{}, f.defaultInformer)
}

func (f *attachmentInformer) Lister() v1alpha1.AttachmentLister {
	return v1alpha1.NewAttachmentLister(f.Informer().GetIndexer())
}
