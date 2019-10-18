// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/minj131/jobcoin/pkg/api"
	"github.com/minj131/jobcoin/pkg/jobcoin"
)

type API struct {
	GetAddressInfoStub        func(string) (*api.AddressInfo, error)
	getAddressInfoMutex       sync.RWMutex
	getAddressInfoArgsForCall []struct {
		arg1 string
	}
	getAddressInfoReturns struct {
		result1 *api.AddressInfo
		result2 error
	}
	getAddressInfoReturnsOnCall map[int]struct {
		result1 *api.AddressInfo
		result2 error
	}
	GetTransactionsStub        func() ([]api.Transactions, error)
	getTransactionsMutex       sync.RWMutex
	getTransactionsArgsForCall []struct {
	}
	getTransactionsReturns struct {
		result1 []api.Transactions
		result2 error
	}
	getTransactionsReturnsOnCall map[int]struct {
		result1 []api.Transactions
		result2 error
	}
	PostTransactionStub        func(string, string, string) error
	postTransactionMutex       sync.RWMutex
	postTransactionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	postTransactionReturns struct {
		result1 error
	}
	postTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *API) GetAddressInfo(arg1 string) (*api.AddressInfo, error) {
	fake.getAddressInfoMutex.Lock()
	ret, specificReturn := fake.getAddressInfoReturnsOnCall[len(fake.getAddressInfoArgsForCall)]
	fake.getAddressInfoArgsForCall = append(fake.getAddressInfoArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetAddressInfo", []interface{}{arg1})
	fake.getAddressInfoMutex.Unlock()
	if fake.GetAddressInfoStub != nil {
		return fake.GetAddressInfoStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAddressInfoReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *API) GetAddressInfoCallCount() int {
	fake.getAddressInfoMutex.RLock()
	defer fake.getAddressInfoMutex.RUnlock()
	return len(fake.getAddressInfoArgsForCall)
}

func (fake *API) GetAddressInfoCalls(stub func(string) (*api.AddressInfo, error)) {
	fake.getAddressInfoMutex.Lock()
	defer fake.getAddressInfoMutex.Unlock()
	fake.GetAddressInfoStub = stub
}

func (fake *API) GetAddressInfoArgsForCall(i int) string {
	fake.getAddressInfoMutex.RLock()
	defer fake.getAddressInfoMutex.RUnlock()
	argsForCall := fake.getAddressInfoArgsForCall[i]
	return argsForCall.arg1
}

func (fake *API) GetAddressInfoReturns(result1 *api.AddressInfo, result2 error) {
	fake.getAddressInfoMutex.Lock()
	defer fake.getAddressInfoMutex.Unlock()
	fake.GetAddressInfoStub = nil
	fake.getAddressInfoReturns = struct {
		result1 *api.AddressInfo
		result2 error
	}{result1, result2}
}

func (fake *API) GetAddressInfoReturnsOnCall(i int, result1 *api.AddressInfo, result2 error) {
	fake.getAddressInfoMutex.Lock()
	defer fake.getAddressInfoMutex.Unlock()
	fake.GetAddressInfoStub = nil
	if fake.getAddressInfoReturnsOnCall == nil {
		fake.getAddressInfoReturnsOnCall = make(map[int]struct {
			result1 *api.AddressInfo
			result2 error
		})
	}
	fake.getAddressInfoReturnsOnCall[i] = struct {
		result1 *api.AddressInfo
		result2 error
	}{result1, result2}
}

func (fake *API) GetTransactions() ([]api.Transactions, error) {
	fake.getTransactionsMutex.Lock()
	ret, specificReturn := fake.getTransactionsReturnsOnCall[len(fake.getTransactionsArgsForCall)]
	fake.getTransactionsArgsForCall = append(fake.getTransactionsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetTransactions", []interface{}{})
	fake.getTransactionsMutex.Unlock()
	if fake.GetTransactionsStub != nil {
		return fake.GetTransactionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getTransactionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *API) GetTransactionsCallCount() int {
	fake.getTransactionsMutex.RLock()
	defer fake.getTransactionsMutex.RUnlock()
	return len(fake.getTransactionsArgsForCall)
}

func (fake *API) GetTransactionsCalls(stub func() ([]api.Transactions, error)) {
	fake.getTransactionsMutex.Lock()
	defer fake.getTransactionsMutex.Unlock()
	fake.GetTransactionsStub = stub
}

func (fake *API) GetTransactionsReturns(result1 []api.Transactions, result2 error) {
	fake.getTransactionsMutex.Lock()
	defer fake.getTransactionsMutex.Unlock()
	fake.GetTransactionsStub = nil
	fake.getTransactionsReturns = struct {
		result1 []api.Transactions
		result2 error
	}{result1, result2}
}

func (fake *API) GetTransactionsReturnsOnCall(i int, result1 []api.Transactions, result2 error) {
	fake.getTransactionsMutex.Lock()
	defer fake.getTransactionsMutex.Unlock()
	fake.GetTransactionsStub = nil
	if fake.getTransactionsReturnsOnCall == nil {
		fake.getTransactionsReturnsOnCall = make(map[int]struct {
			result1 []api.Transactions
			result2 error
		})
	}
	fake.getTransactionsReturnsOnCall[i] = struct {
		result1 []api.Transactions
		result2 error
	}{result1, result2}
}

func (fake *API) PostTransaction(arg1 string, arg2 string, arg3 string) error {
	fake.postTransactionMutex.Lock()
	ret, specificReturn := fake.postTransactionReturnsOnCall[len(fake.postTransactionArgsForCall)]
	fake.postTransactionArgsForCall = append(fake.postTransactionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("PostTransaction", []interface{}{arg1, arg2, arg3})
	fake.postTransactionMutex.Unlock()
	if fake.PostTransactionStub != nil {
		return fake.PostTransactionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.postTransactionReturns
	return fakeReturns.result1
}

func (fake *API) PostTransactionCallCount() int {
	fake.postTransactionMutex.RLock()
	defer fake.postTransactionMutex.RUnlock()
	return len(fake.postTransactionArgsForCall)
}

func (fake *API) PostTransactionCalls(stub func(string, string, string) error) {
	fake.postTransactionMutex.Lock()
	defer fake.postTransactionMutex.Unlock()
	fake.PostTransactionStub = stub
}

func (fake *API) PostTransactionArgsForCall(i int) (string, string, string) {
	fake.postTransactionMutex.RLock()
	defer fake.postTransactionMutex.RUnlock()
	argsForCall := fake.postTransactionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *API) PostTransactionReturns(result1 error) {
	fake.postTransactionMutex.Lock()
	defer fake.postTransactionMutex.Unlock()
	fake.PostTransactionStub = nil
	fake.postTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *API) PostTransactionReturnsOnCall(i int, result1 error) {
	fake.postTransactionMutex.Lock()
	defer fake.postTransactionMutex.Unlock()
	fake.PostTransactionStub = nil
	if fake.postTransactionReturnsOnCall == nil {
		fake.postTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.postTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *API) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAddressInfoMutex.RLock()
	defer fake.getAddressInfoMutex.RUnlock()
	fake.getTransactionsMutex.RLock()
	defer fake.getTransactionsMutex.RUnlock()
	fake.postTransactionMutex.RLock()
	defer fake.postTransactionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *API) recordInvocation(key string, args []interface{}) {
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

var _ jobcoin.API = new(API)
