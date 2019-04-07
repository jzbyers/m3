// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/dbnode/storage/series/buffer.go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package series is a generated GoMock package.
package series

import (
	"reflect"
	"time"

	"github.com/m3db/m3/src/dbnode/storage/block"
	"github.com/m3db/m3/src/dbnode/x/xio"
	"github.com/m3db/m3/src/x/context"
	time0 "github.com/m3db/m3/src/x/time"

	"github.com/golang/mock/gomock"
)

// MockdatabaseBuffer is a mock of databaseBuffer interface
type MockdatabaseBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockdatabaseBufferMockRecorder
}

// MockdatabaseBufferMockRecorder is the mock recorder for MockdatabaseBuffer
type MockdatabaseBufferMockRecorder struct {
	mock *MockdatabaseBuffer
}

// NewMockdatabaseBuffer creates a new mock instance
func NewMockdatabaseBuffer(ctrl *gomock.Controller) *MockdatabaseBuffer {
	mock := &MockdatabaseBuffer{ctrl: ctrl}
	mock.recorder = &MockdatabaseBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockdatabaseBuffer) EXPECT() *MockdatabaseBufferMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockdatabaseBuffer) Write(ctx context.Context, timestamp time.Time, value float64, unit time0.Unit, annotation []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", ctx, timestamp, value, unit, annotation)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockdatabaseBufferMockRecorder) Write(ctx, timestamp, value, unit, annotation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockdatabaseBuffer)(nil).Write), ctx, timestamp, value, unit, annotation)
}

// Snapshot mocks base method
func (m *MockdatabaseBuffer) Snapshot(ctx context.Context, blockStart time.Time) (xio.SegmentReader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", ctx, blockStart)
	ret0, _ := ret[0].(xio.SegmentReader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockdatabaseBufferMockRecorder) Snapshot(ctx, blockStart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockdatabaseBuffer)(nil).Snapshot), ctx, blockStart)
}

// ReadEncoded mocks base method
func (m *MockdatabaseBuffer) ReadEncoded(ctx context.Context, start, end time.Time) [][]xio.BlockReader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadEncoded", ctx, start, end)
	ret0, _ := ret[0].([][]xio.BlockReader)
	return ret0
}

// ReadEncoded indicates an expected call of ReadEncoded
func (mr *MockdatabaseBufferMockRecorder) ReadEncoded(ctx, start, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadEncoded", reflect.TypeOf((*MockdatabaseBuffer)(nil).ReadEncoded), ctx, start, end)
}

// FetchBlocks mocks base method
func (m *MockdatabaseBuffer) FetchBlocks(ctx context.Context, starts []time.Time) []block.FetchBlockResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocks", ctx, starts)
	ret0, _ := ret[0].([]block.FetchBlockResult)
	return ret0
}

// FetchBlocks indicates an expected call of FetchBlocks
func (mr *MockdatabaseBufferMockRecorder) FetchBlocks(ctx, starts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocks", reflect.TypeOf((*MockdatabaseBuffer)(nil).FetchBlocks), ctx, starts)
}

// FetchBlocksMetadata mocks base method
func (m *MockdatabaseBuffer) FetchBlocksMetadata(ctx context.Context, start, end time.Time, opts FetchBlocksMetadataOptions) block.FetchBlockMetadataResults {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlocksMetadata", ctx, start, end, opts)
	ret0, _ := ret[0].(block.FetchBlockMetadataResults)
	return ret0
}

// FetchBlocksMetadata indicates an expected call of FetchBlocksMetadata
func (mr *MockdatabaseBufferMockRecorder) FetchBlocksMetadata(ctx, start, end, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlocksMetadata", reflect.TypeOf((*MockdatabaseBuffer)(nil).FetchBlocksMetadata), ctx, start, end, opts)
}

// IsEmpty mocks base method
func (m *MockdatabaseBuffer) IsEmpty() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEmpty")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEmpty indicates an expected call of IsEmpty
func (mr *MockdatabaseBufferMockRecorder) IsEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEmpty", reflect.TypeOf((*MockdatabaseBuffer)(nil).IsEmpty))
}

// Stats mocks base method
func (m *MockdatabaseBuffer) Stats() bufferStats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(bufferStats)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockdatabaseBufferMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockdatabaseBuffer)(nil).Stats))
}

// MinMax mocks base method
func (m *MockdatabaseBuffer) MinMax() (time.Time, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MinMax")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// MinMax indicates an expected call of MinMax
func (mr *MockdatabaseBufferMockRecorder) MinMax() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MinMax", reflect.TypeOf((*MockdatabaseBuffer)(nil).MinMax))
}

// Tick mocks base method
func (m *MockdatabaseBuffer) Tick() bufferTickResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tick")
	ret0, _ := ret[0].(bufferTickResult)
	return ret0
}

// Tick indicates an expected call of Tick
func (mr *MockdatabaseBufferMockRecorder) Tick() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tick", reflect.TypeOf((*MockdatabaseBuffer)(nil).Tick))
}

// NeedsDrain mocks base method
func (m *MockdatabaseBuffer) NeedsDrain() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsDrain")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedsDrain indicates an expected call of NeedsDrain
func (mr *MockdatabaseBufferMockRecorder) NeedsDrain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsDrain", reflect.TypeOf((*MockdatabaseBuffer)(nil).NeedsDrain))
}

// DrainAndReset mocks base method
func (m *MockdatabaseBuffer) DrainAndReset() drainAndResetResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DrainAndReset")
	ret0, _ := ret[0].(drainAndResetResult)
	return ret0
}

// DrainAndReset indicates an expected call of DrainAndReset
func (mr *MockdatabaseBufferMockRecorder) DrainAndReset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DrainAndReset", reflect.TypeOf((*MockdatabaseBuffer)(nil).DrainAndReset))
}

// Bootstrap mocks base method
func (m *MockdatabaseBuffer) Bootstrap(bl block.DatabaseBlock) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bootstrap", bl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bootstrap indicates an expected call of Bootstrap
func (mr *MockdatabaseBufferMockRecorder) Bootstrap(bl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bootstrap", reflect.TypeOf((*MockdatabaseBuffer)(nil).Bootstrap), bl)
}

// Reset mocks base method
func (m *MockdatabaseBuffer) Reset(opts Options) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset", opts)
}

// Reset indicates an expected call of Reset
func (mr *MockdatabaseBufferMockRecorder) Reset(opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockdatabaseBuffer)(nil).Reset), opts)
}
