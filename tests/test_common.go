package tests

import (
	"github.com/stretchr/testify/mock"
	"message_queue_service/tests/mock_repos"
)

var MockUserRepo = new(mock_repos.MockUserRepo)
var MockProductRepo = new(mock_repos.MockProductRepo)

func InitializeMockFunctions(functions []func() *mock.Call) []*mock.Call {
	var mockCalls = make([]*mock.Call, 0)
	for _, function := range functions {
		mockCall := function()
		mockCalls = append(mockCalls, mockCall)
	}
	return mockCalls
}

func UnsetMockCalls(functions []*mock.Call) {
	for _, call := range functions {
		call.Unset()
	}
}
