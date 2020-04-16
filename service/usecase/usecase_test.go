package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yukinooz/go_task/service/domain"
	"testing"
	"time"
)

type MockTaskRepository struct {
	mock.Mock
}

func (_m *MockTaskRepository) Add(s string, i int, t time.Time) error {
	args := _m.Called(s, i, t)
	return args.Error(0)
}

func (_m *MockTaskRepository) Delete(i int) error {
	args := _m.Called(i)
	return args.Error(0)
}

func (_m *MockTaskRepository) Change(i int, s string, in interface{}) error {
	args := _m.Called(s, i, in)
	return args.Error(0)
}

func (_m *MockTaskRepository) Journal(t time.Time) ([]domain.Task, error) {
	args := _m.Called(t)
	return args.Get(0).([]domain.Task), args.Error(0)
}

func (_m *MockTaskRepository) List(b bool) ([]domain.Task, error) {
	args := _m.Called(b)
	return args.Get(0).([]domain.Task), args.Error(0)
}

var changeCases = []struct {
	msg     string
	id      int
	column  string
	data    string
	wantErr bool
}{
	// deadline
	{msg: "deadline success", id: 5, column: "deadline", data: "3", wantErr: false},

	// name
	{msg: "name success", id: 4, column: "name", data: "new name", wantErr: false},

	// priority
	{msg: "priority changed to high", id: 3, column: "priority", data: "high", wantErr: false},
	{msg: "priority changed to normal", id: 3, column: "priority", data: "normal", wantErr: false},
	{msg: "priority changed to low", id: 3, column: "priority", data: "low", wantErr: false},

	// status
	{msg: "status changed to todo", id: 2, column: "status", data: "todo", wantErr: false},
	{msg: "status changed to doing", id: 2, column: "status", data: "doing", wantErr: false},
	{msg: "status changed to pending", id: 2, column: "status", data: "pending", wantErr: false},
	{msg: "status changed to done", id: 2, column: "status", data: "done", wantErr: false},

	// error
	{msg: "invalid column", id: 2, column: "stats", data: "done", wantErr: true},
	{msg: "invalid data deadline", id: 5, column: "deadline", data: "aa", wantErr: true},
	{msg: "invalid data status", id: 2, column: "status", data: "3", wantErr: true},
	{msg: "invalid data priority", id: 2, column: "priority", data: "do", wantErr: true},
}

func TestChange(t *testing.T) {
	for _, ca := range changeCases {
		repo := &MockTaskRepository{}
		u := NewUsecase(repo)
		err := u.Change(ca.id, ca.column, ca.data)
		if (err != nil) != ca.wantErr {
			t.Errorf("wantErr %t, Err: %s", ca.wantErr, err)
		}
	}
}

//func TestGetDeadlineData(t *testing.T) {
//	cases := []struct {
//		msg string
//		input string
//		want interface{}
//		wantErr bool
//	}{
//		{"invalid input","aa",nil,true},
//		{"input 0","0",24 - time.Now().Hour(),false},
//
//
//	}
//}
