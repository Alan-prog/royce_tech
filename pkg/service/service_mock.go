package service

import (
	"context"
	"github.com/stretchr/testify/mock"
	"my_projects/royce_tech/pkg/models"
)

// MockService ...
type MockService struct {
	mock.Mock
}

// Alive ...
func (m *MockService) Alive(ctx context.Context) (output models.AliveResponse, err error) {
	output.Text = "service is okay"
	return
}

// CreateUser ...
func (m *MockService) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err error) {
	args := m.Called(context.Background(), input)
	if a, ok := args.Get(0).(models.SingleUserData); ok {
		return a, args.Error(1)
	}
	return output, args.Error(1)
}

func (m *MockService) GetUser(ctx context.Context, input int) (output models.SingleUserData, err error) {
	return
}

func (m *MockService) DeleteUser(ctx context.Context, input int) (err error) {
	return
}

func (m *MockService) UpdateUser(ctx context.Context, input *models.UpdateUserData) (output models.SingleUserData, err error) {
	return
}

func (m *MockService) GetAllUser(ctx context.Context) (output models.AllUsersData, err error) {
	return
}
