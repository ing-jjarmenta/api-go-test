package jsonencodec

import "github.com/stretchr/testify/mock"

type MockEncoder struct {
	mock.Mock
}

func (m *MockEncoder) Encode(v any) error {
	args := m.Called(v)

	return args.Error(0)
}
