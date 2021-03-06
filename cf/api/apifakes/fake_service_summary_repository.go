// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeServiceSummaryRepository struct {
	GetSummariesInCurrentSpaceStub        func() ([]models.ServiceInstance, error)
	getSummariesInCurrentSpaceMutex       sync.RWMutex
	getSummariesInCurrentSpaceArgsForCall []struct {
	}
	getSummariesInCurrentSpaceReturns struct {
		result1 []models.ServiceInstance
		result2 error
	}
	getSummariesInCurrentSpaceReturnsOnCall map[int]struct {
		result1 []models.ServiceInstance
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceSummaryRepository) GetSummariesInCurrentSpace() ([]models.ServiceInstance, error) {
	fake.getSummariesInCurrentSpaceMutex.Lock()
	ret, specificReturn := fake.getSummariesInCurrentSpaceReturnsOnCall[len(fake.getSummariesInCurrentSpaceArgsForCall)]
	fake.getSummariesInCurrentSpaceArgsForCall = append(fake.getSummariesInCurrentSpaceArgsForCall, struct {
	}{})
	fake.recordInvocation("GetSummariesInCurrentSpace", []interface{}{})
	fake.getSummariesInCurrentSpaceMutex.Unlock()
	if fake.GetSummariesInCurrentSpaceStub != nil {
		return fake.GetSummariesInCurrentSpaceStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSummariesInCurrentSpaceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceSummaryRepository) GetSummariesInCurrentSpaceCallCount() int {
	fake.getSummariesInCurrentSpaceMutex.RLock()
	defer fake.getSummariesInCurrentSpaceMutex.RUnlock()
	return len(fake.getSummariesInCurrentSpaceArgsForCall)
}

func (fake *FakeServiceSummaryRepository) GetSummariesInCurrentSpaceCalls(stub func() ([]models.ServiceInstance, error)) {
	fake.getSummariesInCurrentSpaceMutex.Lock()
	defer fake.getSummariesInCurrentSpaceMutex.Unlock()
	fake.GetSummariesInCurrentSpaceStub = stub
}

func (fake *FakeServiceSummaryRepository) GetSummariesInCurrentSpaceReturns(result1 []models.ServiceInstance, result2 error) {
	fake.getSummariesInCurrentSpaceMutex.Lock()
	defer fake.getSummariesInCurrentSpaceMutex.Unlock()
	fake.GetSummariesInCurrentSpaceStub = nil
	fake.getSummariesInCurrentSpaceReturns = struct {
		result1 []models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceSummaryRepository) GetSummariesInCurrentSpaceReturnsOnCall(i int, result1 []models.ServiceInstance, result2 error) {
	fake.getSummariesInCurrentSpaceMutex.Lock()
	defer fake.getSummariesInCurrentSpaceMutex.Unlock()
	fake.GetSummariesInCurrentSpaceStub = nil
	if fake.getSummariesInCurrentSpaceReturnsOnCall == nil {
		fake.getSummariesInCurrentSpaceReturnsOnCall = make(map[int]struct {
			result1 []models.ServiceInstance
			result2 error
		})
	}
	fake.getSummariesInCurrentSpaceReturnsOnCall[i] = struct {
		result1 []models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceSummaryRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSummariesInCurrentSpaceMutex.RLock()
	defer fake.getSummariesInCurrentSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceSummaryRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.ServiceSummaryRepository = new(FakeServiceSummaryRepository)
