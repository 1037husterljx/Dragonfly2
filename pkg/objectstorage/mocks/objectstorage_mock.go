// Code generated by MockGen. DO NOT EDIT.
// Source: objectstorage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	objectstorage "d7y.io/dragonfly/v2/pkg/objectstorage"
	gomock "github.com/golang/mock/gomock"
)

// MockObjectStorage is a mock of ObjectStorage interface.
type MockObjectStorage struct {
	ctrl     *gomock.Controller
	recorder *MockObjectStorageMockRecorder
}

// MockObjectStorageMockRecorder is the mock recorder for MockObjectStorage.
type MockObjectStorageMockRecorder struct {
	mock *MockObjectStorage
}

// NewMockObjectStorage creates a new mock instance.
func NewMockObjectStorage(ctrl *gomock.Controller) *MockObjectStorage {
	mock := &MockObjectStorage{ctrl: ctrl}
	mock.recorder = &MockObjectStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectStorage) EXPECT() *MockObjectStorageMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockObjectStorage) CreateBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockObjectStorageMockRecorder) CreateBucket(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockObjectStorage)(nil).CreateBucket), ctx, bucketName)
}

// CreateObject mocks base method.
func (m *MockObjectStorage) CreateObject(ctx context.Context, bucketName, objectKey, digest string, reader io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateObject", ctx, bucketName, objectKey, digest, reader)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateObject indicates an expected call of CreateObject.
func (mr *MockObjectStorageMockRecorder) CreateObject(ctx, bucketName, objectKey, digest, reader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateObject", reflect.TypeOf((*MockObjectStorage)(nil).CreateObject), ctx, bucketName, objectKey, digest, reader)
}

// DeleteBucket mocks base method.
func (m *MockObjectStorage) DeleteBucket(ctx context.Context, bucketName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", ctx, bucketName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockObjectStorageMockRecorder) DeleteBucket(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockObjectStorage)(nil).DeleteBucket), ctx, bucketName)
}

// DeleteObject mocks base method.
func (m *MockObjectStorage) DeleteObject(ctx context.Context, bucketName, objectKey string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObject", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockObjectStorageMockRecorder) DeleteObject(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockObjectStorage)(nil).DeleteObject), ctx, bucketName, objectKey)
}

// GetBucketMetadata mocks base method.
func (m *MockObjectStorage) GetBucketMetadata(ctx context.Context, bucketName string) (*objectstorage.BucketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucketMetadata", ctx, bucketName)
	ret0, _ := ret[0].(*objectstorage.BucketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketMetadata indicates an expected call of GetBucketMetadata.
func (mr *MockObjectStorageMockRecorder) GetBucketMetadata(ctx, bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketMetadata", reflect.TypeOf((*MockObjectStorage)(nil).GetBucketMetadata), ctx, bucketName)
}

// GetObjectMetadata mocks base method.
func (m *MockObjectStorage) GetObjectMetadata(ctx context.Context, bucketName, objectKey string) (*objectstorage.ObjectMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetObjectMetadata", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(*objectstorage.ObjectMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectMetadata indicates an expected call of GetObjectMetadata.
func (mr *MockObjectStorageMockRecorder) GetObjectMetadata(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectMetadata", reflect.TypeOf((*MockObjectStorage)(nil).GetObjectMetadata), ctx, bucketName, objectKey)
}

// GetOject mocks base method.
func (m *MockObjectStorage) GetOject(ctx context.Context, bucketName, objectKey string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOject", ctx, bucketName, objectKey)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOject indicates an expected call of GetOject.
func (mr *MockObjectStorageMockRecorder) GetOject(ctx, bucketName, objectKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOject", reflect.TypeOf((*MockObjectStorage)(nil).GetOject), ctx, bucketName, objectKey)
}

// GetSignURL mocks base method.
func (m *MockObjectStorage) GetSignURL(ctx context.Context, bucketName, objectKey string, method objectstorage.Method, expire time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignURL", ctx, bucketName, objectKey, method, expire)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignURL indicates an expected call of GetSignURL.
func (mr *MockObjectStorageMockRecorder) GetSignURL(ctx, bucketName, objectKey, method, expire interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignURL", reflect.TypeOf((*MockObjectStorage)(nil).GetSignURL), ctx, bucketName, objectKey, method, expire)
}

// ListBucketMetadatas mocks base method.
func (m *MockObjectStorage) ListBucketMetadatas(ctx context.Context) ([]*objectstorage.BucketMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBucketMetadatas", ctx)
	ret0, _ := ret[0].([]*objectstorage.BucketMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketMetadatas indicates an expected call of ListBucketMetadatas.
func (mr *MockObjectStorageMockRecorder) ListBucketMetadatas(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketMetadatas", reflect.TypeOf((*MockObjectStorage)(nil).ListBucketMetadatas), ctx)
}

// ListObjectMetadatas mocks base method.
func (m *MockObjectStorage) ListObjectMetadatas(ctx context.Context, bucketName, prefix, marker string, limit int64) ([]*objectstorage.ObjectMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListObjectMetadatas", ctx, bucketName, prefix, marker, limit)
	ret0, _ := ret[0].([]*objectstorage.ObjectMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectMetadatas indicates an expected call of ListObjectMetadatas.
func (mr *MockObjectStorageMockRecorder) ListObjectMetadatas(ctx, bucketName, prefix, marker, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectMetadatas", reflect.TypeOf((*MockObjectStorage)(nil).ListObjectMetadatas), ctx, bucketName, prefix, marker, limit)
}
