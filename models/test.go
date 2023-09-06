package models

import "github.com/stretchr/testify/mock"

type TestCase struct {
	Case           string
	Input          []interface{}
	MockFunctions  []func() *mock.Call
	ExpectedOutput interface{}
	ExpectedError  error
}
