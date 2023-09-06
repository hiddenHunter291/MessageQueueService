package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"message_queue_service/db"
	"message_queue_service/models"
	"message_queue_service/tests"
	"message_queue_service/tests/test_cases"
	"os"
	"testing"
)

var (
	mockUserRepo    db.UserRepo
	TestUserService UserService
)

func TestMain(m *testing.M) {
	mockUserRepo = tests.MockUserRepo
	TestUserService = NewUserService(mockUserRepo)
	fmt.Println("----------Started test-cases for user----------")
	code := m.Run()
	fmt.Println("----------Completed test-cases for user----------")
	os.Exit(code)
}

func TestGetUser(t *testing.T) {
	for _, test := range test_cases.GetUserTestCases {
		t.Run(test.Case, func(t *testing.T) {
			var userId int
			userId = test.Input[0].(int)

			mockCalls := tests.InitializeMockFunctions(test.MockFunctions)

			actualResult, actualError := TestUserService.Get(userId)
			expectedResult := test.ExpectedOutput.(models.User)
			expectedError := test.ExpectedError
			assert.Equal(t, expectedResult, actualResult, "actual result and expected result are not equal")
			assert.Equal(t, expectedError, actualError, "actual error and expected error are not equal")

			tests.UnsetMockCalls(mockCalls)
		})
	}
}
