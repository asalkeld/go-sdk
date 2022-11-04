// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nitrictech/go-sdk/api/events (interfaces: Events,Topic)

// Package mockapi is a generated GoMock package.
package mockapi

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/nitrictech/apis/go/nitric/v1"
	events "github.com/nitrictech/go-sdk/api/events"
)

// MockEvents is a mock of Events interface.
type MockEvents struct {
	ctrl     *gomock.Controller
	recorder *MockEventsMockRecorder
}

// MockEventsMockRecorder is the mock recorder for MockEvents.
type MockEventsMockRecorder struct {
	mock *MockEvents
}

// NewMockEvents creates a new mock instance.
func NewMockEvents(ctrl *gomock.Controller) *MockEvents {
	mock := &MockEvents{ctrl: ctrl}
	mock.recorder = &MockEventsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvents) EXPECT() *MockEventsMockRecorder {
	return m.recorder
}

// Topic mocks base method.
func (m *MockEvents) Topic(arg0 string) events.Topic {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topic", arg0)
	ret0, _ := ret[0].(events.Topic)
	return ret0
}

// Topic indicates an expected call of Topic.
func (mr *MockEventsMockRecorder) Topic(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topic", reflect.TypeOf((*MockEvents)(nil).Topic), arg0)
}

// Topics mocks base method.
func (m *MockEvents) Topics() ([]events.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Topics")
	ret0, _ := ret[0].([]events.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Topics indicates an expected call of Topics.
func (mr *MockEventsMockRecorder) Topics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Topics", reflect.TypeOf((*MockEvents)(nil).Topics))
}

// MockTopic is a mock of Topic interface.
type MockTopic struct {
	ctrl     *gomock.Controller
	recorder *MockTopicMockRecorder
}

// MockTopicMockRecorder is the mock recorder for MockTopic.
type MockTopicMockRecorder struct {
	mock *MockTopic
}

// NewMockTopic creates a new mock instance.
func NewMockTopic(ctrl *gomock.Controller) *MockTopic {
	mock := &MockTopic{ctrl: ctrl}
	mock.recorder = &MockTopicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTopic) EXPECT() *MockTopicMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockTopic) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockTopicMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockTopic)(nil).Name))
}

// Publish mocks base method.
func (m *MockTopic) Publish(arg0 context.Context, arg1 *events.Event, arg2 ...func(*v1.EventPublishRequest)) (*events.Event, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*events.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockTopicMockRecorder) Publish(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockTopic)(nil).Publish), varargs...)
}
