// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian/storage (interfaces: AdminStorage,AdminTX,LogStorage,LogTreeTX,ReadOnlyAdminTX,ReadOnlyLogTreeTX)

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	trillian "github.com/google/trillian"
	tree "github.com/google/trillian/storage/tree"
	compact "github.com/transparency-dev/merkle/compact"
)

// MockAdminStorage is a mock of AdminStorage interface.
type MockAdminStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAdminStorageMockRecorder
}

// MockAdminStorageMockRecorder is the mock recorder for MockAdminStorage.
type MockAdminStorageMockRecorder struct {
	mock *MockAdminStorage
}

// NewMockAdminStorage creates a new mock instance.
func NewMockAdminStorage(ctrl *gomock.Controller) *MockAdminStorage {
	mock := &MockAdminStorage{ctrl: ctrl}
	mock.recorder = &MockAdminStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminStorage) EXPECT() *MockAdminStorageMockRecorder {
	return m.recorder
}

// CheckDatabaseAccessible mocks base method.
func (m *MockAdminStorage) CheckDatabaseAccessible(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDatabaseAccessible", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDatabaseAccessible indicates an expected call of CheckDatabaseAccessible.
func (mr *MockAdminStorageMockRecorder) CheckDatabaseAccessible(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDatabaseAccessible", reflect.TypeOf((*MockAdminStorage)(nil).CheckDatabaseAccessible), arg0)
}

// ReadWriteTransaction mocks base method.
func (m *MockAdminStorage) ReadWriteTransaction(arg0 context.Context, arg1 AdminTXFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWriteTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadWriteTransaction indicates an expected call of ReadWriteTransaction.
func (mr *MockAdminStorageMockRecorder) ReadWriteTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWriteTransaction", reflect.TypeOf((*MockAdminStorage)(nil).ReadWriteTransaction), arg0, arg1)
}

// Snapshot mocks base method.
func (m *MockAdminStorage) Snapshot(arg0 context.Context) (ReadOnlyAdminTX, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0)
	ret0, _ := ret[0].(ReadOnlyAdminTX)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockAdminStorageMockRecorder) Snapshot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockAdminStorage)(nil).Snapshot), arg0)
}

// MockAdminTX is a mock of AdminTX interface.
type MockAdminTX struct {
	ctrl     *gomock.Controller
	recorder *MockAdminTXMockRecorder
}

// MockAdminTXMockRecorder is the mock recorder for MockAdminTX.
type MockAdminTXMockRecorder struct {
	mock *MockAdminTX
}

// NewMockAdminTX creates a new mock instance.
func NewMockAdminTX(ctrl *gomock.Controller) *MockAdminTX {
	mock := &MockAdminTX{ctrl: ctrl}
	mock.recorder = &MockAdminTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminTX) EXPECT() *MockAdminTXMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAdminTX) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAdminTXMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAdminTX)(nil).Close))
}

// Commit mocks base method.
func (m *MockAdminTX) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockAdminTXMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockAdminTX)(nil).Commit))
}

// CreateTree mocks base method.
func (m *MockAdminTX) CreateTree(arg0 context.Context, arg1 *trillian.Tree) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockAdminTXMockRecorder) CreateTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockAdminTX)(nil).CreateTree), arg0, arg1)
}

// GetTree mocks base method.
func (m *MockAdminTX) GetTree(arg0 context.Context, arg1 int64) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockAdminTXMockRecorder) GetTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockAdminTX)(nil).GetTree), arg0, arg1)
}

// HardDeleteTree mocks base method.
func (m *MockAdminTX) HardDeleteTree(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDeleteTree", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HardDeleteTree indicates an expected call of HardDeleteTree.
func (mr *MockAdminTXMockRecorder) HardDeleteTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDeleteTree", reflect.TypeOf((*MockAdminTX)(nil).HardDeleteTree), arg0, arg1)
}

// ListTrees mocks base method.
func (m *MockAdminTX) ListTrees(arg0 context.Context, arg1 bool) ([]*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrees", arg0, arg1)
	ret0, _ := ret[0].([]*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrees indicates an expected call of ListTrees.
func (mr *MockAdminTXMockRecorder) ListTrees(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrees", reflect.TypeOf((*MockAdminTX)(nil).ListTrees), arg0, arg1)
}

// SoftDeleteTree mocks base method.
func (m *MockAdminTX) SoftDeleteTree(arg0 context.Context, arg1 int64) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SoftDeleteTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SoftDeleteTree indicates an expected call of SoftDeleteTree.
func (mr *MockAdminTXMockRecorder) SoftDeleteTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SoftDeleteTree", reflect.TypeOf((*MockAdminTX)(nil).SoftDeleteTree), arg0, arg1)
}

// UndeleteTree mocks base method.
func (m *MockAdminTX) UndeleteTree(arg0 context.Context, arg1 int64) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndeleteTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UndeleteTree indicates an expected call of UndeleteTree.
func (mr *MockAdminTXMockRecorder) UndeleteTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndeleteTree", reflect.TypeOf((*MockAdminTX)(nil).UndeleteTree), arg0, arg1)
}

// UpdateTree mocks base method.
func (m *MockAdminTX) UpdateTree(arg0 context.Context, arg1 int64, arg2 func(*trillian.Tree)) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTree", arg0, arg1, arg2)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTree indicates an expected call of UpdateTree.
func (mr *MockAdminTXMockRecorder) UpdateTree(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTree", reflect.TypeOf((*MockAdminTX)(nil).UpdateTree), arg0, arg1, arg2)
}

// MockLogStorage is a mock of LogStorage interface.
type MockLogStorage struct {
	ctrl     *gomock.Controller
	recorder *MockLogStorageMockRecorder
}

// MockLogStorageMockRecorder is the mock recorder for MockLogStorage.
type MockLogStorageMockRecorder struct {
	mock *MockLogStorage
}

// NewMockLogStorage creates a new mock instance.
func NewMockLogStorage(ctrl *gomock.Controller) *MockLogStorage {
	mock := &MockLogStorage{ctrl: ctrl}
	mock.recorder = &MockLogStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogStorage) EXPECT() *MockLogStorageMockRecorder {
	return m.recorder
}

// AddSequencedLeaves mocks base method.
func (m *MockLogStorage) AddSequencedLeaves(arg0 context.Context, arg1 *trillian.Tree, arg2 []*trillian.LogLeaf, arg3 time.Time) ([]*trillian.QueuedLogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSequencedLeaves", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*trillian.QueuedLogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSequencedLeaves indicates an expected call of AddSequencedLeaves.
func (mr *MockLogStorageMockRecorder) AddSequencedLeaves(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSequencedLeaves", reflect.TypeOf((*MockLogStorage)(nil).AddSequencedLeaves), arg0, arg1, arg2, arg3)
}

// CheckDatabaseAccessible mocks base method.
func (m *MockLogStorage) CheckDatabaseAccessible(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDatabaseAccessible", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckDatabaseAccessible indicates an expected call of CheckDatabaseAccessible.
func (mr *MockLogStorageMockRecorder) CheckDatabaseAccessible(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDatabaseAccessible", reflect.TypeOf((*MockLogStorage)(nil).CheckDatabaseAccessible), arg0)
}

// GetActiveLogIDs mocks base method.
func (m *MockLogStorage) GetActiveLogIDs(arg0 context.Context) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveLogIDs", arg0)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveLogIDs indicates an expected call of GetActiveLogIDs.
func (mr *MockLogStorageMockRecorder) GetActiveLogIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveLogIDs", reflect.TypeOf((*MockLogStorage)(nil).GetActiveLogIDs), arg0)
}

// QueueLeaves mocks base method.
func (m *MockLogStorage) QueueLeaves(arg0 context.Context, arg1 *trillian.Tree, arg2 []*trillian.LogLeaf, arg3 time.Time) ([]*trillian.QueuedLogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueueLeaves", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*trillian.QueuedLogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueLeaves indicates an expected call of QueueLeaves.
func (mr *MockLogStorageMockRecorder) QueueLeaves(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueLeaves", reflect.TypeOf((*MockLogStorage)(nil).QueueLeaves), arg0, arg1, arg2, arg3)
}

// ReadWriteTransaction mocks base method.
func (m *MockLogStorage) ReadWriteTransaction(arg0 context.Context, arg1 *trillian.Tree, arg2 LogTXFunc) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWriteTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadWriteTransaction indicates an expected call of ReadWriteTransaction.
func (mr *MockLogStorageMockRecorder) ReadWriteTransaction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWriteTransaction", reflect.TypeOf((*MockLogStorage)(nil).ReadWriteTransaction), arg0, arg1, arg2)
}

// SnapshotForTree mocks base method.
func (m *MockLogStorage) SnapshotForTree(arg0 context.Context, arg1 *trillian.Tree) (ReadOnlyLogTreeTX, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotForTree", arg0, arg1)
	ret0, _ := ret[0].(ReadOnlyLogTreeTX)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotForTree indicates an expected call of SnapshotForTree.
func (mr *MockLogStorageMockRecorder) SnapshotForTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotForTree", reflect.TypeOf((*MockLogStorage)(nil).SnapshotForTree), arg0, arg1)
}

// MockLogTreeTX is a mock of LogTreeTX interface.
type MockLogTreeTX struct {
	ctrl     *gomock.Controller
	recorder *MockLogTreeTXMockRecorder
}

// MockLogTreeTXMockRecorder is the mock recorder for MockLogTreeTX.
type MockLogTreeTXMockRecorder struct {
	mock *MockLogTreeTX
}

// NewMockLogTreeTX creates a new mock instance.
func NewMockLogTreeTX(ctrl *gomock.Controller) *MockLogTreeTX {
	mock := &MockLogTreeTX{ctrl: ctrl}
	mock.recorder = &MockLogTreeTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogTreeTX) EXPECT() *MockLogTreeTXMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockLogTreeTX) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockLogTreeTXMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLogTreeTX)(nil).Close))
}

// Commit mocks base method.
func (m *MockLogTreeTX) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockLogTreeTXMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockLogTreeTX)(nil).Commit), arg0)
}

// DequeueLeaves mocks base method.
func (m *MockLogTreeTX) DequeueLeaves(arg0 context.Context, arg1 int, arg2 time.Time) ([]*trillian.LogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DequeueLeaves", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*trillian.LogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DequeueLeaves indicates an expected call of DequeueLeaves.
func (mr *MockLogTreeTXMockRecorder) DequeueLeaves(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DequeueLeaves", reflect.TypeOf((*MockLogTreeTX)(nil).DequeueLeaves), arg0, arg1, arg2)
}

// GetLeavesByHash mocks base method.
func (m *MockLogTreeTX) GetLeavesByHash(arg0 context.Context, arg1 [][]byte, arg2 bool) ([]*trillian.LogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeavesByHash", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*trillian.LogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByHash indicates an expected call of GetLeavesByHash.
func (mr *MockLogTreeTXMockRecorder) GetLeavesByHash(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByHash", reflect.TypeOf((*MockLogTreeTX)(nil).GetLeavesByHash), arg0, arg1, arg2)
}

// GetLeavesByRange mocks base method.
func (m *MockLogTreeTX) GetLeavesByRange(arg0 context.Context, arg1, arg2 int64) ([]*trillian.LogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeavesByRange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*trillian.LogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByRange indicates an expected call of GetLeavesByRange.
func (mr *MockLogTreeTXMockRecorder) GetLeavesByRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByRange", reflect.TypeOf((*MockLogTreeTX)(nil).GetLeavesByRange), arg0, arg1, arg2)
}

// GetMerkleNodes mocks base method.
func (m *MockLogTreeTX) GetMerkleNodes(arg0 context.Context, arg1 []compact.NodeID) ([]tree.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMerkleNodes", arg0, arg1)
	ret0, _ := ret[0].([]tree.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMerkleNodes indicates an expected call of GetMerkleNodes.
func (mr *MockLogTreeTXMockRecorder) GetMerkleNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMerkleNodes", reflect.TypeOf((*MockLogTreeTX)(nil).GetMerkleNodes), arg0, arg1)
}

// LatestSignedLogRoot mocks base method.
func (m *MockLogTreeTX) LatestSignedLogRoot(arg0 context.Context) (*trillian.SignedLogRoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestSignedLogRoot", arg0)
	ret0, _ := ret[0].(*trillian.SignedLogRoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestSignedLogRoot indicates an expected call of LatestSignedLogRoot.
func (mr *MockLogTreeTXMockRecorder) LatestSignedLogRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestSignedLogRoot", reflect.TypeOf((*MockLogTreeTX)(nil).LatestSignedLogRoot), arg0)
}

// LatestStableSignedLogRoot mocks base method.
func (m *MockLogTreeTX) LatestStableSignedLogRoot(arg0 context.Context) (*trillian.SignedLogRoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestStableSignedLogRoot", arg0)
	ret0, _ := ret[0].(*trillian.SignedLogRoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestStableSignedLogRoot indicates an expected call of LatestStableSignedLogRoot.
func (mr *MockLogTreeTXMockRecorder) LatestStableSignedLogRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestStableSignedLogRoot", reflect.TypeOf((*MockLogTreeTX)(nil).LatestStableSignedLogRoot), arg0)
}

// SetMerkleNodes mocks base method.
func (m *MockLogTreeTX) SetMerkleNodes(arg0 context.Context, arg1 []tree.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMerkleNodes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMerkleNodes indicates an expected call of SetMerkleNodes.
func (mr *MockLogTreeTXMockRecorder) SetMerkleNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMerkleNodes", reflect.TypeOf((*MockLogTreeTX)(nil).SetMerkleNodes), arg0, arg1)
}

// StoreSignedLogRoot mocks base method.
func (m *MockLogTreeTX) StoreSignedLogRoot(arg0 context.Context, arg1 *trillian.SignedLogRoot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreSignedLogRoot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreSignedLogRoot indicates an expected call of StoreSignedLogRoot.
func (mr *MockLogTreeTXMockRecorder) StoreSignedLogRoot(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreSignedLogRoot", reflect.TypeOf((*MockLogTreeTX)(nil).StoreSignedLogRoot), arg0, arg1)
}

// UpdateSequencedLeaves mocks base method.
func (m *MockLogTreeTX) UpdateSequencedLeaves(arg0 context.Context, arg1 []*trillian.LogLeaf) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSequencedLeaves", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSequencedLeaves indicates an expected call of UpdateSequencedLeaves.
func (mr *MockLogTreeTXMockRecorder) UpdateSequencedLeaves(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSequencedLeaves", reflect.TypeOf((*MockLogTreeTX)(nil).UpdateSequencedLeaves), arg0, arg1)
}

// MockReadOnlyAdminTX is a mock of ReadOnlyAdminTX interface.
type MockReadOnlyAdminTX struct {
	ctrl     *gomock.Controller
	recorder *MockReadOnlyAdminTXMockRecorder
}

// MockReadOnlyAdminTXMockRecorder is the mock recorder for MockReadOnlyAdminTX.
type MockReadOnlyAdminTXMockRecorder struct {
	mock *MockReadOnlyAdminTX
}

// NewMockReadOnlyAdminTX creates a new mock instance.
func NewMockReadOnlyAdminTX(ctrl *gomock.Controller) *MockReadOnlyAdminTX {
	mock := &MockReadOnlyAdminTX{ctrl: ctrl}
	mock.recorder = &MockReadOnlyAdminTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadOnlyAdminTX) EXPECT() *MockReadOnlyAdminTXMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReadOnlyAdminTX) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReadOnlyAdminTXMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadOnlyAdminTX)(nil).Close))
}

// Commit mocks base method.
func (m *MockReadOnlyAdminTX) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockReadOnlyAdminTXMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockReadOnlyAdminTX)(nil).Commit))
}

// GetTree mocks base method.
func (m *MockReadOnlyAdminTX) GetTree(arg0 context.Context, arg1 int64) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockReadOnlyAdminTXMockRecorder) GetTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockReadOnlyAdminTX)(nil).GetTree), arg0, arg1)
}

// ListTrees mocks base method.
func (m *MockReadOnlyAdminTX) ListTrees(arg0 context.Context, arg1 bool) ([]*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrees", arg0, arg1)
	ret0, _ := ret[0].([]*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrees indicates an expected call of ListTrees.
func (mr *MockReadOnlyAdminTXMockRecorder) ListTrees(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrees", reflect.TypeOf((*MockReadOnlyAdminTX)(nil).ListTrees), arg0, arg1)
}

// MockReadOnlyLogTreeTX is a mock of ReadOnlyLogTreeTX interface.
type MockReadOnlyLogTreeTX struct {
	ctrl     *gomock.Controller
	recorder *MockReadOnlyLogTreeTXMockRecorder
}

// MockReadOnlyLogTreeTXMockRecorder is the mock recorder for MockReadOnlyLogTreeTX.
type MockReadOnlyLogTreeTXMockRecorder struct {
	mock *MockReadOnlyLogTreeTX
}

// NewMockReadOnlyLogTreeTX creates a new mock instance.
func NewMockReadOnlyLogTreeTX(ctrl *gomock.Controller) *MockReadOnlyLogTreeTX {
	mock := &MockReadOnlyLogTreeTX{ctrl: ctrl}
	mock.recorder = &MockReadOnlyLogTreeTXMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadOnlyLogTreeTX) EXPECT() *MockReadOnlyLogTreeTXMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockReadOnlyLogTreeTX) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockReadOnlyLogTreeTXMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).Close))
}

// Commit mocks base method.
func (m *MockReadOnlyLogTreeTX) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockReadOnlyLogTreeTXMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).Commit), arg0)
}

// GetLeavesByHash mocks base method.
func (m *MockReadOnlyLogTreeTX) GetLeavesByHash(arg0 context.Context, arg1 [][]byte, arg2 bool) ([]*trillian.LogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeavesByHash", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*trillian.LogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByHash indicates an expected call of GetLeavesByHash.
func (mr *MockReadOnlyLogTreeTXMockRecorder) GetLeavesByHash(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByHash", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).GetLeavesByHash), arg0, arg1, arg2)
}

// GetLeavesByRange mocks base method.
func (m *MockReadOnlyLogTreeTX) GetLeavesByRange(arg0 context.Context, arg1, arg2 int64) ([]*trillian.LogLeaf, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLeavesByRange", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*trillian.LogLeaf)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLeavesByRange indicates an expected call of GetLeavesByRange.
func (mr *MockReadOnlyLogTreeTXMockRecorder) GetLeavesByRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLeavesByRange", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).GetLeavesByRange), arg0, arg1, arg2)
}

// GetMerkleNodes mocks base method.
func (m *MockReadOnlyLogTreeTX) GetMerkleNodes(arg0 context.Context, arg1 []compact.NodeID) ([]tree.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMerkleNodes", arg0, arg1)
	ret0, _ := ret[0].([]tree.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMerkleNodes indicates an expected call of GetMerkleNodes.
func (mr *MockReadOnlyLogTreeTXMockRecorder) GetMerkleNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMerkleNodes", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).GetMerkleNodes), arg0, arg1)
}

// LatestSignedLogRoot mocks base method.
func (m *MockReadOnlyLogTreeTX) LatestSignedLogRoot(arg0 context.Context) (*trillian.SignedLogRoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestSignedLogRoot", arg0)
	ret0, _ := ret[0].(*trillian.SignedLogRoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestSignedLogRoot indicates an expected call of LatestSignedLogRoot.
func (mr *MockReadOnlyLogTreeTXMockRecorder) LatestSignedLogRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestSignedLogRoot", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).LatestSignedLogRoot), arg0)
}

// LatestStableSignedLogRoot mocks base method.
func (m *MockReadOnlyLogTreeTX) LatestStableSignedLogRoot(arg0 context.Context) (*trillian.SignedLogRoot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestStableSignedLogRoot", arg0)
	ret0, _ := ret[0].(*trillian.SignedLogRoot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestStableSignedLogRoot indicates an expected call of LatestStableSignedLogRoot.
func (mr *MockReadOnlyLogTreeTXMockRecorder) LatestStableSignedLogRoot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestStableSignedLogRoot", reflect.TypeOf((*MockReadOnlyLogTreeTX)(nil).LatestStableSignedLogRoot), arg0)
}
