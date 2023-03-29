// Code generated by counterfeiter. DO NOT EDIT.
package dockerdriverfakes

import (
	"context"
	"sync"

	"code.cloudfoundry.org/dockerdriver"
	"code.cloudfoundry.org/lager/v3"
)

type FakeEnv struct {
	LoggerStub        func() lager.Logger
	loggerMutex       sync.RWMutex
	loggerArgsForCall []struct{}
	loggerReturns     struct {
		result1 lager.Logger
	}
	loggerReturnsOnCall map[int]struct {
		result1 lager.Logger
	}
	ContextStub        func() context.Context
	contextMutex       sync.RWMutex
	contextArgsForCall []struct{}
	contextReturns     struct {
		result1 context.Context
	}
	contextReturnsOnCall map[int]struct {
		result1 context.Context
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnv) Logger() lager.Logger {
	fake.loggerMutex.Lock()
	ret, specificReturn := fake.loggerReturnsOnCall[len(fake.loggerArgsForCall)]
	fake.loggerArgsForCall = append(fake.loggerArgsForCall, struct{}{})
	fake.recordInvocation("Logger", []interface{}{})
	fake.loggerMutex.Unlock()
	if fake.LoggerStub != nil {
		return fake.LoggerStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.loggerReturns.result1
}

func (fake *FakeEnv) LoggerCallCount() int {
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	return len(fake.loggerArgsForCall)
}

func (fake *FakeEnv) LoggerReturns(result1 lager.Logger) {
	fake.LoggerStub = nil
	fake.loggerReturns = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeEnv) LoggerReturnsOnCall(i int, result1 lager.Logger) {
	fake.LoggerStub = nil
	if fake.loggerReturnsOnCall == nil {
		fake.loggerReturnsOnCall = make(map[int]struct {
			result1 lager.Logger
		})
	}
	fake.loggerReturnsOnCall[i] = struct {
		result1 lager.Logger
	}{result1}
}

func (fake *FakeEnv) Context() context.Context {
	fake.contextMutex.Lock()
	ret, specificReturn := fake.contextReturnsOnCall[len(fake.contextArgsForCall)]
	fake.contextArgsForCall = append(fake.contextArgsForCall, struct{}{})
	fake.recordInvocation("Context", []interface{}{})
	fake.contextMutex.Unlock()
	if fake.ContextStub != nil {
		return fake.ContextStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.contextReturns.result1
}

func (fake *FakeEnv) ContextCallCount() int {
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	return len(fake.contextArgsForCall)
}

func (fake *FakeEnv) ContextReturns(result1 context.Context) {
	fake.ContextStub = nil
	fake.contextReturns = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeEnv) ContextReturnsOnCall(i int, result1 context.Context) {
	fake.ContextStub = nil
	if fake.contextReturnsOnCall == nil {
		fake.contextReturnsOnCall = make(map[int]struct {
			result1 context.Context
		})
	}
	fake.contextReturnsOnCall[i] = struct {
		result1 context.Context
	}{result1}
}

func (fake *FakeEnv) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loggerMutex.RLock()
	defer fake.loggerMutex.RUnlock()
	fake.contextMutex.RLock()
	defer fake.contextMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnv) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dockerdriver.Env = new(FakeEnv)
