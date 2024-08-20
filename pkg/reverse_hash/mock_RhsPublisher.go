// Code generated by mockery v2.44.1. DO NOT EDIT.

package reverse_hash

import (
	context "context"

	merkletree_proof "github.com/iden3/merkletree-proof"
	domain "github.com/polygonid/sh-id-platform/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockRhsPublisher is an autogenerated mock type for the RhsPublisher type
type MockRhsPublisher struct {
	mock.Mock
}

// PublishNodesToRHS provides a mock function with given fields: ctx, nodes
func (_m *MockRhsPublisher) PublishNodesToRHS(ctx context.Context, nodes []merkletree_proof.Node) error {
	ret := _m.Called(ctx, nodes)

	if len(ret) == 0 {
		panic("no return value specified for PublishNodesToRHS")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []merkletree_proof.Node) error); ok {
		r0 = rf(ctx, nodes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PushHashesToRHS provides a mock function with given fields: ctx, newState, prevState, revocations, trees
func (_m *MockRhsPublisher) PushHashesToRHS(ctx context.Context, newState *domain.IdentityState, prevState *domain.IdentityState, revocations []*domain.Revocation, trees *domain.IdentityMerkleTrees) error {
	ret := _m.Called(ctx, newState, prevState, revocations, trees)

	if len(ret) == 0 {
		panic("no return value specified for PushHashesToRHS")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.IdentityState, *domain.IdentityState, []*domain.Revocation, *domain.IdentityMerkleTrees) error); ok {
		r0 = rf(ctx, newState, prevState, revocations, trees)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRhsPublisher creates a new instance of MockRhsPublisher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRhsPublisher(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRhsPublisher {
	mock := &MockRhsPublisher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
