package mock_repos

import (
	"github.com/stretchr/testify/mock"
	"message_queue_service/models"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) CreateProduct(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepo) CreateOrder(userID int, productName, productDescription string) error {
	args := m.Called(userID, productName, productDescription)
	return args.Error(0)
}

func (m *MockProductRepo) SaveLocalPath(productID int, path string) error {
	args := m.Called(productID, path)
	return args.Error(0)
}

func (m *MockProductRepo) GetProductURLs(productID int) ([]string, error) {
	args := m.Called(productID)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockProductRepo) GetProductID(productName, productDescription string) (int, error) {
	args := m.Called(productName, productDescription)
	return args.Int(0), args.Error(1)
}
