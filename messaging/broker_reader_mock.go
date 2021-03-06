// Code generated by MockGen. DO NOT EDIT.
// Source: reader.go

// Package messaging is a generated GoMock package.
package messaging

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReader) Read(msgFunc ReadMessageFunc) error {
	ret := m.ctrl.Call(m, "Read", msgFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// Read indicates an expected call of Read
func (mr *MockReaderMockRecorder) Read(msgFunc interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReader)(nil).Read), msgFunc)
}

// Close mocks base method
func (m *MockReader) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockReaderMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReader)(nil).Close))
}

// MockBrokerReader is a mock of BrokerReader interface
type MockBrokerReader struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerReaderMockRecorder
}

// MockBrokerReaderMockRecorder is the mock recorder for MockBrokerReader
type MockBrokerReaderMockRecorder struct {
	mock *MockBrokerReader
}

// NewMockBrokerReader creates a new mock instance
func NewMockBrokerReader(ctrl *gomock.Controller) *MockBrokerReader {
	mock := &MockBrokerReader{ctrl: ctrl}
	mock.recorder = &MockBrokerReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBrokerReader) EXPECT() *MockBrokerReaderMockRecorder {
	return m.recorder
}

// FetchMessage mocks base method
func (m *MockBrokerReader) FetchMessage(ctx context.Context) (Message, error) {
	ret := m.ctrl.Call(m, "FetchMessage", ctx)
	ret0, _ := ret[0].(Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMessage indicates an expected call of FetchMessage
func (mr *MockBrokerReaderMockRecorder) FetchMessage(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMessage", reflect.TypeOf((*MockBrokerReader)(nil).FetchMessage), ctx)
}

// CommitMessages mocks base method
func (m *MockBrokerReader) CommitMessages(ctx context.Context, msgs ...Message) error {
	varargs := []interface{}{ctx}
	for _, a := range msgs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CommitMessages", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitMessages indicates an expected call of CommitMessages
func (mr *MockBrokerReaderMockRecorder) CommitMessages(ctx interface{}, msgs ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx}, msgs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitMessages", reflect.TypeOf((*MockBrokerReader)(nil).CommitMessages), varargs...)
}

// ReadMessage mocks base method
func (m *MockBrokerReader) ReadMessage(ctx context.Context) (Message, error) {
	ret := m.ctrl.Call(m, "ReadMessage", ctx)
	ret0, _ := ret[0].(Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessage indicates an expected call of ReadMessage
func (mr *MockBrokerReaderMockRecorder) ReadMessage(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessage", reflect.TypeOf((*MockBrokerReader)(nil).ReadMessage), ctx)
}

// Close mocks base method
func (m *MockBrokerReader) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockBrokerReaderMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBrokerReader)(nil).Close))
}
