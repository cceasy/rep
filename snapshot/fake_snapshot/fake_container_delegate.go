// This file was generated by counterfeiter
package fake_snapshot

import (
	"sync"

	"github.com/cloudfoundry-incubator/rep/snapshot"
	"github.com/pivotal-golang/lager"
)

type FakeContainerDelegate struct {
	RunContainerStub        func(logger lager.Logger, guid string) bool
	runContainerMutex       sync.RWMutex
	runContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	runContainerReturns struct {
		result1 bool
	}
	StopContainerStub        func(logger lager.Logger, guid string) bool
	stopContainerMutex       sync.RWMutex
	stopContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	stopContainerReturns struct {
		result1 bool
	}
	DeleteContainerStub        func(logger lager.Logger, guid string) bool
	deleteContainerMutex       sync.RWMutex
	deleteContainerArgsForCall []struct {
		logger lager.Logger
		guid   string
	}
	deleteContainerReturns struct {
		result1 bool
	}
	FetchContainerResultStub        func(logger lager.Logger, guid string, filename string) (string, error)
	fetchContainerResultMutex       sync.RWMutex
	fetchContainerResultArgsForCall []struct {
		logger   lager.Logger
		guid     string
		filename string
	}
	fetchContainerResultReturns struct {
		result1 string
		result2 error
	}
}

func (fake *FakeContainerDelegate) RunContainer(logger lager.Logger, guid string) bool {
	fake.runContainerMutex.Lock()
	fake.runContainerArgsForCall = append(fake.runContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.runContainerMutex.Unlock()
	if fake.RunContainerStub != nil {
		return fake.RunContainerStub(logger, guid)
	} else {
		return fake.runContainerReturns.result1
	}
}

func (fake *FakeContainerDelegate) RunContainerCallCount() int {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return len(fake.runContainerArgsForCall)
}

func (fake *FakeContainerDelegate) RunContainerArgsForCall(i int) (lager.Logger, string) {
	fake.runContainerMutex.RLock()
	defer fake.runContainerMutex.RUnlock()
	return fake.runContainerArgsForCall[i].logger, fake.runContainerArgsForCall[i].guid
}

func (fake *FakeContainerDelegate) RunContainerReturns(result1 bool) {
	fake.RunContainerStub = nil
	fake.runContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerDelegate) StopContainer(logger lager.Logger, guid string) bool {
	fake.stopContainerMutex.Lock()
	fake.stopContainerArgsForCall = append(fake.stopContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.stopContainerMutex.Unlock()
	if fake.StopContainerStub != nil {
		return fake.StopContainerStub(logger, guid)
	} else {
		return fake.stopContainerReturns.result1
	}
}

func (fake *FakeContainerDelegate) StopContainerCallCount() int {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return len(fake.stopContainerArgsForCall)
}

func (fake *FakeContainerDelegate) StopContainerArgsForCall(i int) (lager.Logger, string) {
	fake.stopContainerMutex.RLock()
	defer fake.stopContainerMutex.RUnlock()
	return fake.stopContainerArgsForCall[i].logger, fake.stopContainerArgsForCall[i].guid
}

func (fake *FakeContainerDelegate) StopContainerReturns(result1 bool) {
	fake.StopContainerStub = nil
	fake.stopContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerDelegate) DeleteContainer(logger lager.Logger, guid string) bool {
	fake.deleteContainerMutex.Lock()
	fake.deleteContainerArgsForCall = append(fake.deleteContainerArgsForCall, struct {
		logger lager.Logger
		guid   string
	}{logger, guid})
	fake.deleteContainerMutex.Unlock()
	if fake.DeleteContainerStub != nil {
		return fake.DeleteContainerStub(logger, guid)
	} else {
		return fake.deleteContainerReturns.result1
	}
}

func (fake *FakeContainerDelegate) DeleteContainerCallCount() int {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return len(fake.deleteContainerArgsForCall)
}

func (fake *FakeContainerDelegate) DeleteContainerArgsForCall(i int) (lager.Logger, string) {
	fake.deleteContainerMutex.RLock()
	defer fake.deleteContainerMutex.RUnlock()
	return fake.deleteContainerArgsForCall[i].logger, fake.deleteContainerArgsForCall[i].guid
}

func (fake *FakeContainerDelegate) DeleteContainerReturns(result1 bool) {
	fake.DeleteContainerStub = nil
	fake.deleteContainerReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeContainerDelegate) FetchContainerResult(logger lager.Logger, guid string, filename string) (string, error) {
	fake.fetchContainerResultMutex.Lock()
	fake.fetchContainerResultArgsForCall = append(fake.fetchContainerResultArgsForCall, struct {
		logger   lager.Logger
		guid     string
		filename string
	}{logger, guid, filename})
	fake.fetchContainerResultMutex.Unlock()
	if fake.FetchContainerResultStub != nil {
		return fake.FetchContainerResultStub(logger, guid, filename)
	} else {
		return fake.fetchContainerResultReturns.result1, fake.fetchContainerResultReturns.result2
	}
}

func (fake *FakeContainerDelegate) FetchContainerResultCallCount() int {
	fake.fetchContainerResultMutex.RLock()
	defer fake.fetchContainerResultMutex.RUnlock()
	return len(fake.fetchContainerResultArgsForCall)
}

func (fake *FakeContainerDelegate) FetchContainerResultArgsForCall(i int) (lager.Logger, string, string) {
	fake.fetchContainerResultMutex.RLock()
	defer fake.fetchContainerResultMutex.RUnlock()
	return fake.fetchContainerResultArgsForCall[i].logger, fake.fetchContainerResultArgsForCall[i].guid, fake.fetchContainerResultArgsForCall[i].filename
}

func (fake *FakeContainerDelegate) FetchContainerResultReturns(result1 string, result2 error) {
	fake.FetchContainerResultStub = nil
	fake.fetchContainerResultReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ snapshot.ContainerDelegate = new(FakeContainerDelegate)