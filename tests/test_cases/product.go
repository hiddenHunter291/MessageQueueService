package test_cases

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"message_queue_service/models"
	"message_queue_service/tests"
	"time"
)

var CreateTestCases = []models.TestCase{
	{
		Case:  "Create product (success)",
		Input: []interface{}{models.OrderInfo{}},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateProduct", models.Product{ProductLinks: make([]models.ProductLinks, 0), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}).
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateOrder", 0, "", "").
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("GetProductID", "", "").
					Return(0, nil)
			},
			func() *mock.Call {
				return tests.MockProducer.
					On("ProduceEvent", 0).
					Return(nil)
			},
		},
		ExpectedOutput: nil,
		ExpectedError:  nil,
	},
	{
		Case:  "Create product (product failure)",
		Input: []interface{}{models.OrderInfo{}},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateProduct", models.Product{ProductLinks: make([]models.ProductLinks, 0), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}).
					Return(errors.New("create product failed"))
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateOrder", 0, "", "").
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("GetProductID", "", "").
					Return(0, nil)
			},
			func() *mock.Call {
				return tests.MockProducer.
					On("ProduceEvent", 0).
					Return(nil)
			},
		},
		ExpectedOutput: errors.New("create product failed"),
		ExpectedError:  nil,
	},
	{
		Case:  "Create product (order failure)",
		Input: []interface{}{models.OrderInfo{}},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateProduct", models.Product{ProductLinks: make([]models.ProductLinks, 0), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}).
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateOrder", 0, "", "").
					Return(errors.New("create order failed"))
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("GetProductID", "", "").
					Return(0, nil)
			},
			func() *mock.Call {
				return tests.MockProducer.
					On("ProduceEvent", 0).
					Return(nil)
			},
		},
		ExpectedOutput: errors.New("create order failed"),
		ExpectedError:  nil,
	},
	{
		Case:  "Create product (get productID failure)",
		Input: []interface{}{models.OrderInfo{}},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateProduct", models.Product{ProductLinks: make([]models.ProductLinks, 0), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}).
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateOrder", 0, "", "").
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("GetProductID", "", "").
					Return(0, errors.New("get product Id failed"))
			},
			func() *mock.Call {
				return tests.MockProducer.
					On("ProduceEvent", 0).
					Return(nil)
			},
		},
		ExpectedOutput: errors.New("get product Id failed"),
		ExpectedError:  nil,
	},
	{
		Case:  "Create product (produce event failure)",
		Input: []interface{}{models.OrderInfo{}},
		MockFunctions: []func() *mock.Call{
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateProduct", models.Product{ProductLinks: make([]models.ProductLinks, 0), CreatedAt: time.Now().Unix(), UpdatedAt: time.Now().Unix()}).
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("CreateOrder", 0, "", "").
					Return(nil)
			},
			func() *mock.Call {
				return tests.MockProductRepo.
					On("GetProductID", "", "").
					Return(0, nil)
			},
			func() *mock.Call {
				return tests.MockProducer.
					On("ProduceEvent", 0).
					Return(errors.New("produce event failed"))
			},
		},
		ExpectedOutput: errors.New("produce event failed"),
		ExpectedError:  nil,
	},
}
