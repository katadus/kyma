// Code generated by mockery v1.0.0
package automock

import internal "github.com/kyma-project/kyma/components/helm-broker/internal"
import mock "github.com/stretchr/testify/mock"

// InstanceBindDataRemover is an autogenerated mock type for the InstanceBindDataRemover type
type InstanceBindDataRemover struct {
	mock.Mock
}

// Remove provides a mock function with given fields: _a0
func (_m *InstanceBindDataRemover) Remove(_a0 internal.InstanceID) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(internal.InstanceID) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}