// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/kiln/builder"
)

type Interpolator struct {
	InterpolateStub        func(input builder.InterpolateInput, templateYAML []byte) ([]byte, error)
	interpolateMutex       sync.RWMutex
	interpolateArgsForCall []struct {
		input        builder.InterpolateInput
		templateYAML []byte
	}
	interpolateReturns struct {
		result1 []byte
		result2 error
	}
	interpolateReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Interpolator) Interpolate(input builder.InterpolateInput, templateYAML []byte) ([]byte, error) {
	var templateYAMLCopy []byte
	if templateYAML != nil {
		templateYAMLCopy = make([]byte, len(templateYAML))
		copy(templateYAMLCopy, templateYAML)
	}
	fake.interpolateMutex.Lock()
	ret, specificReturn := fake.interpolateReturnsOnCall[len(fake.interpolateArgsForCall)]
	fake.interpolateArgsForCall = append(fake.interpolateArgsForCall, struct {
		input        builder.InterpolateInput
		templateYAML []byte
	}{input, templateYAMLCopy})
	fake.recordInvocation("Interpolate", []interface{}{input, templateYAMLCopy})
	fake.interpolateMutex.Unlock()
	if fake.InterpolateStub != nil {
		return fake.InterpolateStub(input, templateYAML)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interpolateReturns.result1, fake.interpolateReturns.result2
}

func (fake *Interpolator) InterpolateCallCount() int {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return len(fake.interpolateArgsForCall)
}

func (fake *Interpolator) InterpolateArgsForCall(i int) (builder.InterpolateInput, []byte) {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return fake.interpolateArgsForCall[i].input, fake.interpolateArgsForCall[i].templateYAML
}

func (fake *Interpolator) InterpolateReturns(result1 []byte, result2 error) {
	fake.InterpolateStub = nil
	fake.interpolateReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Interpolator) InterpolateReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.InterpolateStub = nil
	if fake.interpolateReturnsOnCall == nil {
		fake.interpolateReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.interpolateReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *Interpolator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Interpolator) recordInvocation(key string, args []interface{}) {
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
