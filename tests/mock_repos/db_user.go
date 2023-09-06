package mock_repos

import (
	"github.com/stretchr/testify/mock"
	"message_queue_service/models"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) GetUserId(userId int) (models.User, error) {
	args := m.Called(userId)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepo) Set(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}
