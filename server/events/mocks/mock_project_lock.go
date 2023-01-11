// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: ProjectLocker)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
	models "github.com/runatlantis/atlantis/server/events/models"
	logging "github.com/runatlantis/atlantis/server/logging"
	"reflect"
	"time"
)

type MockProjectLocker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockProjectLocker(options ...pegomock.Option) *MockProjectLocker {
	mock := &MockProjectLocker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockProjectLocker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockProjectLocker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockProjectLocker) TryLock(_param0 logging.SimpleLogging, _param1 models.PullRequest, _param2 models.User, _param3 string, _param4 models.Project, _param5 bool) (*events.TryLockResponse, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockProjectLocker().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((**events.TryLockResponse)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *events.TryLockResponse
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*events.TryLockResponse)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockProjectLocker) VerifyWasCalledOnce() *VerifierMockProjectLocker {
	return &VerifierMockProjectLocker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockProjectLocker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockProjectLocker {
	return &VerifierMockProjectLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockProjectLocker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockProjectLocker {
	return &VerifierMockProjectLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockProjectLocker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockProjectLocker {
	return &VerifierMockProjectLocker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockProjectLocker struct {
	mock                   *MockProjectLocker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockProjectLocker) TryLock(_param0 logging.SimpleLogging, _param1 models.PullRequest, _param2 models.User, _param3 string, _param4 models.Project, _param5 bool) *MockProjectLocker_TryLock_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4, _param5}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &MockProjectLocker_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockProjectLocker_TryLock_OngoingVerification struct {
	mock              *MockProjectLocker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockProjectLocker_TryLock_OngoingVerification) GetCapturedArguments() (logging.SimpleLogging, models.PullRequest, models.User, string, models.Project, bool) {
	_param0, _param1, _param2, _param3, _param4, _param5 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1], _param4[len(_param4)-1], _param5[len(_param5)-1]
}

func (c *MockProjectLocker_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []logging.SimpleLogging, _param1 []models.PullRequest, _param2 []models.User, _param3 []string, _param4 []models.Project, _param5 []bool) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]logging.SimpleLogging, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(logging.SimpleLogging)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.User, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(models.User)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]models.Project, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(models.Project)
		}
		_param5 = make([]bool, len(c.methodInvocations))
		for u, param := range params[5] {
			_param5[u] = param.(bool)
		}
	}
	return
}
