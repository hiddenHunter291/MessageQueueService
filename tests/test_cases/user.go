package test_cases

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"message_queue_service/models"
	"message_queue_service/tests"
)

var GetUserTestCases = []models.TestCase{
	{
		Case:  "Get user (success)",
		Input: []interface{}{12},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockUserRepo.
					On("GetUserId", 12).
					Return(models.User{}, nil)
			},
		},
		ExpectedOutput: models.User{},
		ExpectedError:  nil,
	},
	{
		Case:  "Get user (failure)",
		Input: []interface{}{12},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockUserRepo.
					On("GetUserId", 12).
					Return(models.User{}, errors.New("fetch failed"))
			},
		},
		ExpectedOutput: models.User{},
		ExpectedError:  errors.New("fetch failed"),
	},
}
