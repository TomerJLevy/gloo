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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualHostOptionLister helps list VirtualHostOptions.
// All objects returned here must be treated as read-only.
type VirtualHostOptionLister interface {
	// List lists all VirtualHostOptions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualHostOption, err error)
	// VirtualHostOptions returns an object that can list and get VirtualHostOptions.
	VirtualHostOptions(namespace string) VirtualHostOptionNamespaceLister
	VirtualHostOptionListerExpansion
}

// virtualHostOptionLister implements the VirtualHostOptionLister interface.
type virtualHostOptionLister struct {
	indexer cache.Indexer
}

// NewVirtualHostOptionLister returns a new VirtualHostOptionLister.
func NewVirtualHostOptionLister(indexer cache.Indexer) VirtualHostOptionLister {
	return &virtualHostOptionLister{indexer: indexer}
}

// List lists all VirtualHostOptions in the indexer.
func (s *virtualHostOptionLister) List(selector labels.Selector) (ret []*v1.VirtualHostOption, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualHostOption))
	})
	return ret, err
}

// VirtualHostOptions returns an object that can list and get VirtualHostOptions.
func (s *virtualHostOptionLister) VirtualHostOptions(namespace string) VirtualHostOptionNamespaceLister {
	return virtualHostOptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualHostOptionNamespaceLister helps list and get VirtualHostOptions.
// All objects returned here must be treated as read-only.
type VirtualHostOptionNamespaceLister interface {
	// List lists all VirtualHostOptions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualHostOption, err error)
	// Get retrieves the VirtualHostOption from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualHostOption, error)
	VirtualHostOptionNamespaceListerExpansion
}

// virtualHostOptionNamespaceLister implements the VirtualHostOptionNamespaceLister
// interface.
type virtualHostOptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualHostOptions in the indexer for a given namespace.
func (s virtualHostOptionNamespaceLister) List(selector labels.Selector) (ret []*v1.VirtualHostOption, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualHostOption))
	})
	return ret, err
}

// Get retrieves the VirtualHostOption from the indexer for a given namespace and name.
func (s virtualHostOptionNamespaceLister) Get(name string) (*v1.VirtualHostOption, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualhostoption"), name)
	}
	return obj.(*v1.VirtualHostOption), nil
}