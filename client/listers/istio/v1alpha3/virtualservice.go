/*
Copyright 2019 The Knative Authors

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

package v1alpha3

import (
	v1alpha3 "/apis/istio/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualServiceLister helps list VirtualServices.
type VirtualServiceLister interface {
	// List lists all VirtualServices in the indexer.
	List(selector labels.Selector) (ret []*v1alpha3.VirtualService, err error)
	// VirtualServices returns an object that can list and get VirtualServices.
	VirtualServices(namespace string) VirtualServiceNamespaceLister
	VirtualServiceListerExpansion
}

// virtualServiceLister implements the VirtualServiceLister interface.
type virtualServiceLister struct {
	indexer cache.Indexer
}

// NewVirtualServiceLister returns a new VirtualServiceLister.
func NewVirtualServiceLister(indexer cache.Indexer) VirtualServiceLister {
	return &virtualServiceLister{indexer: indexer}
}

// List lists all VirtualServices in the indexer.
func (s *virtualServiceLister) List(selector labels.Selector) (ret []*v1alpha3.VirtualService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.VirtualService))
	})
	return ret, err
}

// VirtualServices returns an object that can list and get VirtualServices.
func (s *virtualServiceLister) VirtualServices(namespace string) VirtualServiceNamespaceLister {
	return virtualServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualServiceNamespaceLister helps list and get VirtualServices.
type VirtualServiceNamespaceLister interface {
	// List lists all VirtualServices in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha3.VirtualService, err error)
	// Get retrieves the VirtualService from the indexer for a given namespace and name.
	Get(name string) (*v1alpha3.VirtualService, error)
	VirtualServiceNamespaceListerExpansion
}

// virtualServiceNamespaceLister implements the VirtualServiceNamespaceLister
// interface.
type virtualServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualServices in the indexer for a given namespace.
func (s virtualServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.VirtualService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.VirtualService))
	})
	return ret, err
}

// Get retrieves the VirtualService from the indexer for a given namespace and name.
func (s virtualServiceNamespaceLister) Get(name string) (*v1alpha3.VirtualService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("virtualservice"), name)
	}
	return obj.(*v1alpha3.VirtualService), nil
}
