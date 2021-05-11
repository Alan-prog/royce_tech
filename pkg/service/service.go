package service

import (
	"context"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
)

type royce interface {
	Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output int, err tools.ErrorMessage)
	GetSingleUser(ctx context.Context, id int) (output models.SingleUserData, err tools.ErrorMessage)
	DeleteUser(ctx context.Context, id int) (err tools.ErrorMessage)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (err tools.ErrorMessage)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err tools.ErrorMessage)
}

type Service interface {
	Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err tools.ErrorMessage)
	GetUser(ctx context.Context, input int) (output models.SingleUserData, err tools.ErrorMessage)
	DeleteUser(ctx context.Context, input int) (err tools.ErrorMessage)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (output models.SingleUserData, err tools.ErrorMessage)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err tools.ErrorMessage)
}

type service struct {
	royce royce
}

func (s *service) Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage) {
	output, err = s.royce.Alive(ctx)
	return
}

func (s *service) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err tools.ErrorMessage) {
	newUserId, err := s.royce.CreateUser(ctx, input)
	if err != nil {
		return
	}

	output, err = s.royce.GetSingleUser(ctx, newUserId)
	return
}

func (s *service) GetUser(ctx context.Context, input int) (output models.SingleUserData, err tools.ErrorMessage) {
	output, err = s.royce.GetSingleUser(ctx, input)
	return
}

func (s *service) DeleteUser(ctx context.Context, id int) (err tools.ErrorMessage) {
	err = s.royce.DeleteUser(ctx, id)
	return
}

func (s *service) UpdateUser(ctx context.Context, input *models.UpdateUserData) (response models.SingleUserData, err tools.ErrorMessage) {
	err = s.royce.UpdateUser(ctx, input)
	if err != nil {
		return
	}

	response, err = s.royce.GetSingleUser(ctx, input.ID)
	return
}

func (s *service) GetAllUser(ctx context.Context) (response models.AllUsersData, err tools.ErrorMessage) {
	response, err = s.royce.GetAllUser(ctx)

	return
}

// NewService ...
func NewService(royce royce) Service {
	return &service{
		royce: royce,
	}
}
