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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/spacesbucket/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ObjectLister helps list Objects.
// All objects returned here must be treated as read-only.
type ObjectLister interface {
	// List lists all Objects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Object, err error)
	// Objects returns an object that can list and get Objects.
	Objects(namespace string) ObjectNamespaceLister
	ObjectListerExpansion
}

// objectLister implements the ObjectLister interface.
type objectLister struct {
	indexer cache.Indexer
}

// NewObjectLister returns a new ObjectLister.
func NewObjectLister(indexer cache.Indexer) ObjectLister {
	return &objectLister{indexer: indexer}
}

// List lists all Objects in the indexer.
func (s *objectLister) List(selector labels.Selector) (ret []*v1alpha1.Object, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Object))
	})
	return ret, err
}

// Objects returns an object that can list and get Objects.
func (s *objectLister) Objects(namespace string) ObjectNamespaceLister {
	return objectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ObjectNamespaceLister helps list and get Objects.
// All objects returned here must be treated as read-only.
type ObjectNamespaceLister interface {
	// List lists all Objects in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Object, err error)
	// Get retrieves the Object from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Object, error)
	ObjectNamespaceListerExpansion
}

// objectNamespaceLister implements the ObjectNamespaceLister
// interface.
type objectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Objects in the indexer for a given namespace.
func (s objectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Object, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Object))
	})
	return ret, err
}

// Get retrieves the Object from the indexer for a given namespace and name.
func (s objectNamespaceLister) Get(name string) (*v1alpha1.Object, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("object"), name)
	}
	return obj.(*v1alpha1.Object), nil
}
