// Code generated by counterfeiter. DO NOT EDIT.
package azurefilebrokerfakes

import (
	"sync"

	"code.cloudfoundry.org/goshims/sqlshim"
	"github.com/AbelHu/azurefilebroker/azurefilebroker"
)

type FakeSqlVariant struct {
	ConnectStub        func() (sqlshim.SqlDB, error)
	connectMutex       sync.RWMutex
	connectArgsForCall []struct{}
	connectReturns     struct {
		result1 sqlshim.SqlDB
		result2 error
	}
	connectReturnsOnCall map[int]struct {
		result1 sqlshim.SqlDB
		result2 error
	}
	GetInitializeDatabaseSQLStub        func() []string
	getInitializeDatabaseSQLMutex       sync.RWMutex
	getInitializeDatabaseSQLArgsForCall []struct{}
	getInitializeDatabaseSQLReturns     struct {
		result1 []string
	}
	getInitializeDatabaseSQLReturnsOnCall map[int]struct {
		result1 []string
	}
	GetAppLockSQLStub        func() string
	getAppLockSQLMutex       sync.RWMutex
	getAppLockSQLArgsForCall []struct{}
	getAppLockSQLReturns     struct {
		result1 string
	}
	getAppLockSQLReturnsOnCall map[int]struct {
		result1 string
	}
	GetReleaseAppLockSQLStub        func() string
	getReleaseAppLockSQLMutex       sync.RWMutex
	getReleaseAppLockSQLArgsForCall []struct{}
	getReleaseAppLockSQLReturns     struct {
		result1 string
	}
	getReleaseAppLockSQLReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSqlVariant) Connect() (sqlshim.SqlDB, error) {
	fake.connectMutex.Lock()
	ret, specificReturn := fake.connectReturnsOnCall[len(fake.connectArgsForCall)]
	fake.connectArgsForCall = append(fake.connectArgsForCall, struct{}{})
	fake.recordInvocation("Connect", []interface{}{})
	fake.connectMutex.Unlock()
	if fake.ConnectStub != nil {
		return fake.ConnectStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.connectReturns.result1, fake.connectReturns.result2
}

func (fake *FakeSqlVariant) ConnectCallCount() int {
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	return len(fake.connectArgsForCall)
}

func (fake *FakeSqlVariant) ConnectReturns(result1 sqlshim.SqlDB, result2 error) {
	fake.ConnectStub = nil
	fake.connectReturns = struct {
		result1 sqlshim.SqlDB
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlVariant) ConnectReturnsOnCall(i int, result1 sqlshim.SqlDB, result2 error) {
	fake.ConnectStub = nil
	if fake.connectReturnsOnCall == nil {
		fake.connectReturnsOnCall = make(map[int]struct {
			result1 sqlshim.SqlDB
			result2 error
		})
	}
	fake.connectReturnsOnCall[i] = struct {
		result1 sqlshim.SqlDB
		result2 error
	}{result1, result2}
}

func (fake *FakeSqlVariant) GetInitializeDatabaseSQL() []string {
	fake.getInitializeDatabaseSQLMutex.Lock()
	ret, specificReturn := fake.getInitializeDatabaseSQLReturnsOnCall[len(fake.getInitializeDatabaseSQLArgsForCall)]
	fake.getInitializeDatabaseSQLArgsForCall = append(fake.getInitializeDatabaseSQLArgsForCall, struct{}{})
	fake.recordInvocation("GetInitializeDatabaseSQL", []interface{}{})
	fake.getInitializeDatabaseSQLMutex.Unlock()
	if fake.GetInitializeDatabaseSQLStub != nil {
		return fake.GetInitializeDatabaseSQLStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getInitializeDatabaseSQLReturns.result1
}

func (fake *FakeSqlVariant) GetInitializeDatabaseSQLCallCount() int {
	fake.getInitializeDatabaseSQLMutex.RLock()
	defer fake.getInitializeDatabaseSQLMutex.RUnlock()
	return len(fake.getInitializeDatabaseSQLArgsForCall)
}

func (fake *FakeSqlVariant) GetInitializeDatabaseSQLReturns(result1 []string) {
	fake.GetInitializeDatabaseSQLStub = nil
	fake.getInitializeDatabaseSQLReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeSqlVariant) GetInitializeDatabaseSQLReturnsOnCall(i int, result1 []string) {
	fake.GetInitializeDatabaseSQLStub = nil
	if fake.getInitializeDatabaseSQLReturnsOnCall == nil {
		fake.getInitializeDatabaseSQLReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getInitializeDatabaseSQLReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeSqlVariant) GetAppLockSQL() string {
	fake.getAppLockSQLMutex.Lock()
	ret, specificReturn := fake.getAppLockSQLReturnsOnCall[len(fake.getAppLockSQLArgsForCall)]
	fake.getAppLockSQLArgsForCall = append(fake.getAppLockSQLArgsForCall, struct{}{})
	fake.recordInvocation("GetAppLockSQL", []interface{}{})
	fake.getAppLockSQLMutex.Unlock()
	if fake.GetAppLockSQLStub != nil {
		return fake.GetAppLockSQLStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getAppLockSQLReturns.result1
}

func (fake *FakeSqlVariant) GetAppLockSQLCallCount() int {
	fake.getAppLockSQLMutex.RLock()
	defer fake.getAppLockSQLMutex.RUnlock()
	return len(fake.getAppLockSQLArgsForCall)
}

func (fake *FakeSqlVariant) GetAppLockSQLReturns(result1 string) {
	fake.GetAppLockSQLStub = nil
	fake.getAppLockSQLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSqlVariant) GetAppLockSQLReturnsOnCall(i int, result1 string) {
	fake.GetAppLockSQLStub = nil
	if fake.getAppLockSQLReturnsOnCall == nil {
		fake.getAppLockSQLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getAppLockSQLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSqlVariant) GetReleaseAppLockSQL() string {
	fake.getReleaseAppLockSQLMutex.Lock()
	ret, specificReturn := fake.getReleaseAppLockSQLReturnsOnCall[len(fake.getReleaseAppLockSQLArgsForCall)]
	fake.getReleaseAppLockSQLArgsForCall = append(fake.getReleaseAppLockSQLArgsForCall, struct{}{})
	fake.recordInvocation("GetReleaseAppLockSQL", []interface{}{})
	fake.getReleaseAppLockSQLMutex.Unlock()
	if fake.GetReleaseAppLockSQLStub != nil {
		return fake.GetReleaseAppLockSQLStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getReleaseAppLockSQLReturns.result1
}

func (fake *FakeSqlVariant) GetReleaseAppLockSQLCallCount() int {
	fake.getReleaseAppLockSQLMutex.RLock()
	defer fake.getReleaseAppLockSQLMutex.RUnlock()
	return len(fake.getReleaseAppLockSQLArgsForCall)
}

func (fake *FakeSqlVariant) GetReleaseAppLockSQLReturns(result1 string) {
	fake.GetReleaseAppLockSQLStub = nil
	fake.getReleaseAppLockSQLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSqlVariant) GetReleaseAppLockSQLReturnsOnCall(i int, result1 string) {
	fake.GetReleaseAppLockSQLStub = nil
	if fake.getReleaseAppLockSQLReturnsOnCall == nil {
		fake.getReleaseAppLockSQLReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getReleaseAppLockSQLReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSqlVariant) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.connectMutex.RLock()
	defer fake.connectMutex.RUnlock()
	fake.getInitializeDatabaseSQLMutex.RLock()
	defer fake.getInitializeDatabaseSQLMutex.RUnlock()
	fake.getAppLockSQLMutex.RLock()
	defer fake.getAppLockSQLMutex.RUnlock()
	fake.getReleaseAppLockSQLMutex.RLock()
	defer fake.getReleaseAppLockSQLMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSqlVariant) recordInvocation(key string, args []interface{}) {
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

var _ azurefilebroker.SqlVariant = new(FakeSqlVariant)
