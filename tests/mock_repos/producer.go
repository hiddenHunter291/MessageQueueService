package mock_repos

import "github.com/stretchr/testify/mock"

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) ProduceEvent(key int) error {
	args := m.Called(key)
	return args.Error(0)
}
