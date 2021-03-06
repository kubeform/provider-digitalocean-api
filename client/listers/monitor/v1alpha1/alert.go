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
	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/monitor/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertLister helps list Alerts.
// All objects returned here must be treated as read-only.
type AlertLister interface {
	// List lists all Alerts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Alert, err error)
	// Alerts returns an object that can list and get Alerts.
	Alerts(namespace string) AlertNamespaceLister
	AlertListerExpansion
}

// alertLister implements the AlertLister interface.
type alertLister struct {
	indexer cache.Indexer
}

// NewAlertLister returns a new AlertLister.
func NewAlertLister(indexer cache.Indexer) AlertLister {
	return &alertLister{indexer: indexer}
}

// List lists all Alerts in the indexer.
func (s *alertLister) List(selector labels.Selector) (ret []*v1alpha1.Alert, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alert))
	})
	return ret, err
}

// Alerts returns an object that can list and get Alerts.
func (s *alertLister) Alerts(namespace string) AlertNamespaceLister {
	return alertNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlertNamespaceLister helps list and get Alerts.
// All objects returned here must be treated as read-only.
type AlertNamespaceLister interface {
	// List lists all Alerts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Alert, err error)
	// Get retrieves the Alert from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Alert, error)
	AlertNamespaceListerExpansion
}

// alertNamespaceLister implements the AlertNamespaceLister
// interface.
type alertNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Alerts in the indexer for a given namespace.
func (s alertNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Alert, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alert))
	})
	return ret, err
}

// Get retrieves the Alert from the indexer for a given namespace and name.
func (s alertNamespaceLister) Get(name string) (*v1alpha1.Alert, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alert"), name)
	}
	return obj.(*v1alpha1.Alert), nil
}
