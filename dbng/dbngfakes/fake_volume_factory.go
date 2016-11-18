// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
)

type FakeVolumeFactory struct {
	GetTeamVolumesStub        func(teamID int) ([]dbng.CreatedVolume, error)
	getTeamVolumesMutex       sync.RWMutex
	getTeamVolumesArgsForCall []struct {
		teamID int
	}
	getTeamVolumesReturns struct {
		result1 []dbng.CreatedVolume
		result2 error
	}
	CreateContainerVolumeStub        func(*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) (dbng.CreatingVolume, error)
	createContainerVolumeMutex       sync.RWMutex
	createContainerVolumeArgsForCall []struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}
	createContainerVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	FindContainerVolumeStub        func(*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) (dbng.CreatingVolume, dbng.CreatedVolume, error)
	findContainerVolumeMutex       sync.RWMutex
	findContainerVolumeArgsForCall []struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}
	findContainerVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}
	FindBaseResourceTypeVolumeStub        func(*dbng.Team, *dbng.Worker, *dbng.UsedBaseResourceType) (dbng.CreatingVolume, dbng.CreatedVolume, error)
	findBaseResourceTypeVolumeMutex       sync.RWMutex
	findBaseResourceTypeVolumeArgsForCall []struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.UsedBaseResourceType
	}
	findBaseResourceTypeVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}
	CreateBaseResourceTypeVolumeStub        func(*dbng.Team, *dbng.Worker, *dbng.UsedBaseResourceType) (dbng.CreatingVolume, error)
	createBaseResourceTypeVolumeMutex       sync.RWMutex
	createBaseResourceTypeVolumeArgsForCall []struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.UsedBaseResourceType
	}
	createBaseResourceTypeVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	FindResourceCacheVolumeStub        func(*dbng.Worker, *dbng.UsedResourceCache) (dbng.CreatingVolume, dbng.CreatedVolume, error)
	findResourceCacheVolumeMutex       sync.RWMutex
	findResourceCacheVolumeArgsForCall []struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
	}
	findResourceCacheVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}
	FindResourceCacheInitializedVolumeStub        func(*dbng.Worker, *dbng.UsedResourceCache) (dbng.CreatedVolume, bool, error)
	findResourceCacheInitializedVolumeMutex       sync.RWMutex
	findResourceCacheInitializedVolumeArgsForCall []struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
	}
	findResourceCacheInitializedVolumeReturns struct {
		result1 dbng.CreatedVolume
		result2 bool
		result3 error
	}
	CreateResourceCacheVolumeStub        func(*dbng.Worker, *dbng.UsedResourceCache) (dbng.CreatingVolume, error)
	createResourceCacheVolumeMutex       sync.RWMutex
	createResourceCacheVolumeArgsForCall []struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
	}
	createResourceCacheVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	CreateResourceCacheVolumeWithParentStub        func(*dbng.Worker, *dbng.UsedResourceCache, string) (dbng.CreatingVolume, error)
	createResourceCacheVolumeWithParentMutex       sync.RWMutex
	createResourceCacheVolumeWithParentArgsForCall []struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
		arg3 string
	}
	createResourceCacheVolumeWithParentReturns struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	FindVolumesForContainerStub        func(containerID int) ([]dbng.CreatedVolume, error)
	findVolumesForContainerMutex       sync.RWMutex
	findVolumesForContainerArgsForCall []struct {
		containerID int
	}
	findVolumesForContainerReturns struct {
		result1 []dbng.CreatedVolume
		result2 error
	}
	GetOrphanedVolumesStub        func() ([]dbng.CreatedVolume, []dbng.DestroyingVolume, error)
	getOrphanedVolumesMutex       sync.RWMutex
	getOrphanedVolumesArgsForCall []struct{}
	getOrphanedVolumesReturns     struct {
		result1 []dbng.CreatedVolume
		result2 []dbng.DestroyingVolume
		result3 error
	}
	FindCreatedVolumeStub        func(handle string) (dbng.CreatedVolume, bool, error)
	findCreatedVolumeMutex       sync.RWMutex
	findCreatedVolumeArgsForCall []struct {
		handle string
	}
	findCreatedVolumeReturns struct {
		result1 dbng.CreatedVolume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeFactory) GetTeamVolumes(teamID int) ([]dbng.CreatedVolume, error) {
	fake.getTeamVolumesMutex.Lock()
	fake.getTeamVolumesArgsForCall = append(fake.getTeamVolumesArgsForCall, struct {
		teamID int
	}{teamID})
	fake.recordInvocation("GetTeamVolumes", []interface{}{teamID})
	fake.getTeamVolumesMutex.Unlock()
	if fake.GetTeamVolumesStub != nil {
		return fake.GetTeamVolumesStub(teamID)
	} else {
		return fake.getTeamVolumesReturns.result1, fake.getTeamVolumesReturns.result2
	}
}

func (fake *FakeVolumeFactory) GetTeamVolumesCallCount() int {
	fake.getTeamVolumesMutex.RLock()
	defer fake.getTeamVolumesMutex.RUnlock()
	return len(fake.getTeamVolumesArgsForCall)
}

func (fake *FakeVolumeFactory) GetTeamVolumesArgsForCall(i int) int {
	fake.getTeamVolumesMutex.RLock()
	defer fake.getTeamVolumesMutex.RUnlock()
	return fake.getTeamVolumesArgsForCall[i].teamID
}

func (fake *FakeVolumeFactory) GetTeamVolumesReturns(result1 []dbng.CreatedVolume, result2 error) {
	fake.GetTeamVolumesStub = nil
	fake.getTeamVolumesReturns = struct {
		result1 []dbng.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateContainerVolume(arg1 *dbng.Team, arg2 *dbng.Worker, arg3 *dbng.CreatingContainer, arg4 string) (dbng.CreatingVolume, error) {
	fake.createContainerVolumeMutex.Lock()
	fake.createContainerVolumeArgsForCall = append(fake.createContainerVolumeArgsForCall, struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateContainerVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.createContainerVolumeMutex.Unlock()
	if fake.CreateContainerVolumeStub != nil {
		return fake.CreateContainerVolumeStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.createContainerVolumeReturns.result1, fake.createContainerVolumeReturns.result2
	}
}

func (fake *FakeVolumeFactory) CreateContainerVolumeCallCount() int {
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	return len(fake.createContainerVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateContainerVolumeArgsForCall(i int) (*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) {
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	return fake.createContainerVolumeArgsForCall[i].arg1, fake.createContainerVolumeArgsForCall[i].arg2, fake.createContainerVolumeArgsForCall[i].arg3, fake.createContainerVolumeArgsForCall[i].arg4
}

func (fake *FakeVolumeFactory) CreateContainerVolumeReturns(result1 dbng.CreatingVolume, result2 error) {
	fake.CreateContainerVolumeStub = nil
	fake.createContainerVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindContainerVolume(arg1 *dbng.Team, arg2 *dbng.Worker, arg3 *dbng.CreatingContainer, arg4 string) (dbng.CreatingVolume, dbng.CreatedVolume, error) {
	fake.findContainerVolumeMutex.Lock()
	fake.findContainerVolumeArgsForCall = append(fake.findContainerVolumeArgsForCall, struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("FindContainerVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.findContainerVolumeMutex.Unlock()
	if fake.FindContainerVolumeStub != nil {
		return fake.FindContainerVolumeStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.findContainerVolumeReturns.result1, fake.findContainerVolumeReturns.result2, fake.findContainerVolumeReturns.result3
	}
}

func (fake *FakeVolumeFactory) FindContainerVolumeCallCount() int {
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return len(fake.findContainerVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindContainerVolumeArgsForCall(i int) (*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) {
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return fake.findContainerVolumeArgsForCall[i].arg1, fake.findContainerVolumeArgsForCall[i].arg2, fake.findContainerVolumeArgsForCall[i].arg3, fake.findContainerVolumeArgsForCall[i].arg4
}

func (fake *FakeVolumeFactory) FindContainerVolumeReturns(result1 dbng.CreatingVolume, result2 dbng.CreatedVolume, result3 error) {
	fake.FindContainerVolumeStub = nil
	fake.findContainerVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolume(arg1 *dbng.Team, arg2 *dbng.Worker, arg3 *dbng.UsedBaseResourceType) (dbng.CreatingVolume, dbng.CreatedVolume, error) {
	fake.findBaseResourceTypeVolumeMutex.Lock()
	fake.findBaseResourceTypeVolumeArgsForCall = append(fake.findBaseResourceTypeVolumeArgsForCall, struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.UsedBaseResourceType
	}{arg1, arg2, arg3})
	fake.recordInvocation("FindBaseResourceTypeVolume", []interface{}{arg1, arg2, arg3})
	fake.findBaseResourceTypeVolumeMutex.Unlock()
	if fake.FindBaseResourceTypeVolumeStub != nil {
		return fake.FindBaseResourceTypeVolumeStub(arg1, arg2, arg3)
	} else {
		return fake.findBaseResourceTypeVolumeReturns.result1, fake.findBaseResourceTypeVolumeReturns.result2, fake.findBaseResourceTypeVolumeReturns.result3
	}
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeCallCount() int {
	fake.findBaseResourceTypeVolumeMutex.RLock()
	defer fake.findBaseResourceTypeVolumeMutex.RUnlock()
	return len(fake.findBaseResourceTypeVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeArgsForCall(i int) (*dbng.Team, *dbng.Worker, *dbng.UsedBaseResourceType) {
	fake.findBaseResourceTypeVolumeMutex.RLock()
	defer fake.findBaseResourceTypeVolumeMutex.RUnlock()
	return fake.findBaseResourceTypeVolumeArgsForCall[i].arg1, fake.findBaseResourceTypeVolumeArgsForCall[i].arg2, fake.findBaseResourceTypeVolumeArgsForCall[i].arg3
}

func (fake *FakeVolumeFactory) FindBaseResourceTypeVolumeReturns(result1 dbng.CreatingVolume, result2 dbng.CreatedVolume, result3 error) {
	fake.FindBaseResourceTypeVolumeStub = nil
	fake.findBaseResourceTypeVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolume(arg1 *dbng.Team, arg2 *dbng.Worker, arg3 *dbng.UsedBaseResourceType) (dbng.CreatingVolume, error) {
	fake.createBaseResourceTypeVolumeMutex.Lock()
	fake.createBaseResourceTypeVolumeArgsForCall = append(fake.createBaseResourceTypeVolumeArgsForCall, struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.UsedBaseResourceType
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateBaseResourceTypeVolume", []interface{}{arg1, arg2, arg3})
	fake.createBaseResourceTypeVolumeMutex.Unlock()
	if fake.CreateBaseResourceTypeVolumeStub != nil {
		return fake.CreateBaseResourceTypeVolumeStub(arg1, arg2, arg3)
	} else {
		return fake.createBaseResourceTypeVolumeReturns.result1, fake.createBaseResourceTypeVolumeReturns.result2
	}
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeCallCount() int {
	fake.createBaseResourceTypeVolumeMutex.RLock()
	defer fake.createBaseResourceTypeVolumeMutex.RUnlock()
	return len(fake.createBaseResourceTypeVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeArgsForCall(i int) (*dbng.Team, *dbng.Worker, *dbng.UsedBaseResourceType) {
	fake.createBaseResourceTypeVolumeMutex.RLock()
	defer fake.createBaseResourceTypeVolumeMutex.RUnlock()
	return fake.createBaseResourceTypeVolumeArgsForCall[i].arg1, fake.createBaseResourceTypeVolumeArgsForCall[i].arg2, fake.createBaseResourceTypeVolumeArgsForCall[i].arg3
}

func (fake *FakeVolumeFactory) CreateBaseResourceTypeVolumeReturns(result1 dbng.CreatingVolume, result2 error) {
	fake.CreateBaseResourceTypeVolumeStub = nil
	fake.createBaseResourceTypeVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindResourceCacheVolume(arg1 *dbng.Worker, arg2 *dbng.UsedResourceCache) (dbng.CreatingVolume, dbng.CreatedVolume, error) {
	fake.findResourceCacheVolumeMutex.Lock()
	fake.findResourceCacheVolumeArgsForCall = append(fake.findResourceCacheVolumeArgsForCall, struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
	}{arg1, arg2})
	fake.recordInvocation("FindResourceCacheVolume", []interface{}{arg1, arg2})
	fake.findResourceCacheVolumeMutex.Unlock()
	if fake.FindResourceCacheVolumeStub != nil {
		return fake.FindResourceCacheVolumeStub(arg1, arg2)
	} else {
		return fake.findResourceCacheVolumeReturns.result1, fake.findResourceCacheVolumeReturns.result2, fake.findResourceCacheVolumeReturns.result3
	}
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeCallCount() int {
	fake.findResourceCacheVolumeMutex.RLock()
	defer fake.findResourceCacheVolumeMutex.RUnlock()
	return len(fake.findResourceCacheVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeArgsForCall(i int) (*dbng.Worker, *dbng.UsedResourceCache) {
	fake.findResourceCacheVolumeMutex.RLock()
	defer fake.findResourceCacheVolumeMutex.RUnlock()
	return fake.findResourceCacheVolumeArgsForCall[i].arg1, fake.findResourceCacheVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) FindResourceCacheVolumeReturns(result1 dbng.CreatingVolume, result2 dbng.CreatedVolume, result3 error) {
	fake.FindResourceCacheVolumeStub = nil
	fake.findResourceCacheVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindResourceCacheInitializedVolume(arg1 *dbng.Worker, arg2 *dbng.UsedResourceCache) (dbng.CreatedVolume, bool, error) {
	fake.findResourceCacheInitializedVolumeMutex.Lock()
	fake.findResourceCacheInitializedVolumeArgsForCall = append(fake.findResourceCacheInitializedVolumeArgsForCall, struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
	}{arg1, arg2})
	fake.recordInvocation("FindResourceCacheInitializedVolume", []interface{}{arg1, arg2})
	fake.findResourceCacheInitializedVolumeMutex.Unlock()
	if fake.FindResourceCacheInitializedVolumeStub != nil {
		return fake.FindResourceCacheInitializedVolumeStub(arg1, arg2)
	} else {
		return fake.findResourceCacheInitializedVolumeReturns.result1, fake.findResourceCacheInitializedVolumeReturns.result2, fake.findResourceCacheInitializedVolumeReturns.result3
	}
}

func (fake *FakeVolumeFactory) FindResourceCacheInitializedVolumeCallCount() int {
	fake.findResourceCacheInitializedVolumeMutex.RLock()
	defer fake.findResourceCacheInitializedVolumeMutex.RUnlock()
	return len(fake.findResourceCacheInitializedVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindResourceCacheInitializedVolumeArgsForCall(i int) (*dbng.Worker, *dbng.UsedResourceCache) {
	fake.findResourceCacheInitializedVolumeMutex.RLock()
	defer fake.findResourceCacheInitializedVolumeMutex.RUnlock()
	return fake.findResourceCacheInitializedVolumeArgsForCall[i].arg1, fake.findResourceCacheInitializedVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) FindResourceCacheInitializedVolumeReturns(result1 dbng.CreatedVolume, result2 bool, result3 error) {
	fake.FindResourceCacheInitializedVolumeStub = nil
	fake.findResourceCacheInitializedVolumeReturns = struct {
		result1 dbng.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolume(arg1 *dbng.Worker, arg2 *dbng.UsedResourceCache) (dbng.CreatingVolume, error) {
	fake.createResourceCacheVolumeMutex.Lock()
	fake.createResourceCacheVolumeArgsForCall = append(fake.createResourceCacheVolumeArgsForCall, struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
	}{arg1, arg2})
	fake.recordInvocation("CreateResourceCacheVolume", []interface{}{arg1, arg2})
	fake.createResourceCacheVolumeMutex.Unlock()
	if fake.CreateResourceCacheVolumeStub != nil {
		return fake.CreateResourceCacheVolumeStub(arg1, arg2)
	} else {
		return fake.createResourceCacheVolumeReturns.result1, fake.createResourceCacheVolumeReturns.result2
	}
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeCallCount() int {
	fake.createResourceCacheVolumeMutex.RLock()
	defer fake.createResourceCacheVolumeMutex.RUnlock()
	return len(fake.createResourceCacheVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeArgsForCall(i int) (*dbng.Worker, *dbng.UsedResourceCache) {
	fake.createResourceCacheVolumeMutex.RLock()
	defer fake.createResourceCacheVolumeMutex.RUnlock()
	return fake.createResourceCacheVolumeArgsForCall[i].arg1, fake.createResourceCacheVolumeArgsForCall[i].arg2
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeReturns(result1 dbng.CreatingVolume, result2 error) {
	fake.CreateResourceCacheVolumeStub = nil
	fake.createResourceCacheVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeWithParent(arg1 *dbng.Worker, arg2 *dbng.UsedResourceCache, arg3 string) (dbng.CreatingVolume, error) {
	fake.createResourceCacheVolumeWithParentMutex.Lock()
	fake.createResourceCacheVolumeWithParentArgsForCall = append(fake.createResourceCacheVolumeWithParentArgsForCall, struct {
		arg1 *dbng.Worker
		arg2 *dbng.UsedResourceCache
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateResourceCacheVolumeWithParent", []interface{}{arg1, arg2, arg3})
	fake.createResourceCacheVolumeWithParentMutex.Unlock()
	if fake.CreateResourceCacheVolumeWithParentStub != nil {
		return fake.CreateResourceCacheVolumeWithParentStub(arg1, arg2, arg3)
	} else {
		return fake.createResourceCacheVolumeWithParentReturns.result1, fake.createResourceCacheVolumeWithParentReturns.result2
	}
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeWithParentCallCount() int {
	fake.createResourceCacheVolumeWithParentMutex.RLock()
	defer fake.createResourceCacheVolumeWithParentMutex.RUnlock()
	return len(fake.createResourceCacheVolumeWithParentArgsForCall)
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeWithParentArgsForCall(i int) (*dbng.Worker, *dbng.UsedResourceCache, string) {
	fake.createResourceCacheVolumeWithParentMutex.RLock()
	defer fake.createResourceCacheVolumeWithParentMutex.RUnlock()
	return fake.createResourceCacheVolumeWithParentArgsForCall[i].arg1, fake.createResourceCacheVolumeWithParentArgsForCall[i].arg2, fake.createResourceCacheVolumeWithParentArgsForCall[i].arg3
}

func (fake *FakeVolumeFactory) CreateResourceCacheVolumeWithParentReturns(result1 dbng.CreatingVolume, result2 error) {
	fake.CreateResourceCacheVolumeWithParentStub = nil
	fake.createResourceCacheVolumeWithParentReturns = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) FindVolumesForContainer(containerID int) ([]dbng.CreatedVolume, error) {
	fake.findVolumesForContainerMutex.Lock()
	fake.findVolumesForContainerArgsForCall = append(fake.findVolumesForContainerArgsForCall, struct {
		containerID int
	}{containerID})
	fake.recordInvocation("FindVolumesForContainer", []interface{}{containerID})
	fake.findVolumesForContainerMutex.Unlock()
	if fake.FindVolumesForContainerStub != nil {
		return fake.FindVolumesForContainerStub(containerID)
	} else {
		return fake.findVolumesForContainerReturns.result1, fake.findVolumesForContainerReturns.result2
	}
}

func (fake *FakeVolumeFactory) FindVolumesForContainerCallCount() int {
	fake.findVolumesForContainerMutex.RLock()
	defer fake.findVolumesForContainerMutex.RUnlock()
	return len(fake.findVolumesForContainerArgsForCall)
}

func (fake *FakeVolumeFactory) FindVolumesForContainerArgsForCall(i int) int {
	fake.findVolumesForContainerMutex.RLock()
	defer fake.findVolumesForContainerMutex.RUnlock()
	return fake.findVolumesForContainerArgsForCall[i].containerID
}

func (fake *FakeVolumeFactory) FindVolumesForContainerReturns(result1 []dbng.CreatedVolume, result2 error) {
	fake.FindVolumesForContainerStub = nil
	fake.findVolumesForContainerReturns = struct {
		result1 []dbng.CreatedVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeFactory) GetOrphanedVolumes() ([]dbng.CreatedVolume, []dbng.DestroyingVolume, error) {
	fake.getOrphanedVolumesMutex.Lock()
	fake.getOrphanedVolumesArgsForCall = append(fake.getOrphanedVolumesArgsForCall, struct{}{})
	fake.recordInvocation("GetOrphanedVolumes", []interface{}{})
	fake.getOrphanedVolumesMutex.Unlock()
	if fake.GetOrphanedVolumesStub != nil {
		return fake.GetOrphanedVolumesStub()
	} else {
		return fake.getOrphanedVolumesReturns.result1, fake.getOrphanedVolumesReturns.result2, fake.getOrphanedVolumesReturns.result3
	}
}

func (fake *FakeVolumeFactory) GetOrphanedVolumesCallCount() int {
	fake.getOrphanedVolumesMutex.RLock()
	defer fake.getOrphanedVolumesMutex.RUnlock()
	return len(fake.getOrphanedVolumesArgsForCall)
}

func (fake *FakeVolumeFactory) GetOrphanedVolumesReturns(result1 []dbng.CreatedVolume, result2 []dbng.DestroyingVolume, result3 error) {
	fake.GetOrphanedVolumesStub = nil
	fake.getOrphanedVolumesReturns = struct {
		result1 []dbng.CreatedVolume
		result2 []dbng.DestroyingVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) FindCreatedVolume(handle string) (dbng.CreatedVolume, bool, error) {
	fake.findCreatedVolumeMutex.Lock()
	fake.findCreatedVolumeArgsForCall = append(fake.findCreatedVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.recordInvocation("FindCreatedVolume", []interface{}{handle})
	fake.findCreatedVolumeMutex.Unlock()
	if fake.FindCreatedVolumeStub != nil {
		return fake.FindCreatedVolumeStub(handle)
	} else {
		return fake.findCreatedVolumeReturns.result1, fake.findCreatedVolumeReturns.result2, fake.findCreatedVolumeReturns.result3
	}
}

func (fake *FakeVolumeFactory) FindCreatedVolumeCallCount() int {
	fake.findCreatedVolumeMutex.RLock()
	defer fake.findCreatedVolumeMutex.RUnlock()
	return len(fake.findCreatedVolumeArgsForCall)
}

func (fake *FakeVolumeFactory) FindCreatedVolumeArgsForCall(i int) string {
	fake.findCreatedVolumeMutex.RLock()
	defer fake.findCreatedVolumeMutex.RUnlock()
	return fake.findCreatedVolumeArgsForCall[i].handle
}

func (fake *FakeVolumeFactory) FindCreatedVolumeReturns(result1 dbng.CreatedVolume, result2 bool, result3 error) {
	fake.FindCreatedVolumeStub = nil
	fake.findCreatedVolumeReturns = struct {
		result1 dbng.CreatedVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVolumeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamVolumesMutex.RLock()
	defer fake.getTeamVolumesMutex.RUnlock()
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	fake.findBaseResourceTypeVolumeMutex.RLock()
	defer fake.findBaseResourceTypeVolumeMutex.RUnlock()
	fake.createBaseResourceTypeVolumeMutex.RLock()
	defer fake.createBaseResourceTypeVolumeMutex.RUnlock()
	fake.findResourceCacheVolumeMutex.RLock()
	defer fake.findResourceCacheVolumeMutex.RUnlock()
	fake.findResourceCacheInitializedVolumeMutex.RLock()
	defer fake.findResourceCacheInitializedVolumeMutex.RUnlock()
	fake.createResourceCacheVolumeMutex.RLock()
	defer fake.createResourceCacheVolumeMutex.RUnlock()
	fake.createResourceCacheVolumeWithParentMutex.RLock()
	defer fake.createResourceCacheVolumeWithParentMutex.RUnlock()
	fake.findVolumesForContainerMutex.RLock()
	defer fake.findVolumesForContainerMutex.RUnlock()
	fake.getOrphanedVolumesMutex.RLock()
	defer fake.getOrphanedVolumesMutex.RUnlock()
	fake.findCreatedVolumeMutex.RLock()
	defer fake.findCreatedVolumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumeFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.VolumeFactory = new(FakeVolumeFactory)
