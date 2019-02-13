// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import v1alpha1 "github.com/kyma-project/kyma/components/assetstore-controller-manager/pkg/apis/assetstore/v1alpha1"

// Lister is an autogenerated mock type for the Lister type
type Lister struct {
	mock.Mock
}

// Get provides a mock function with given fields: namespace, name
func (_m *Lister) Get(namespace string, name string) (*v1alpha1.Bucket, error) {
	ret := _m.Called(namespace, name)

	var r0 *v1alpha1.Bucket
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.Bucket); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Bucket)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
