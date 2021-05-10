package service

import (
	"context"
	"my_projects/royce_tech/pkg/models"
)

type royce interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output int, err error)
	GetSingleUser(ctx context.Context, id int) (output models.SingleUserData, err error)
	DeleteUser(ctx context.Context, id int) (err error)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (err error)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err error)
}

type Service interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err error)
	GetUser(ctx context.Context, input int) (output models.SingleUserData, err error)
	DeleteUser(ctx context.Context, input int) (err error)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (output models.SingleUserData, err error)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err error)
}

type service struct {
	royce royce
}

func (s *service) Alive() (output models.AliveResponse, err error) {
	output, err = s.royce.Alive()
	return
}

func (s *service) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err error) {
	newUserId, err := s.royce.CreateUser(ctx, input)
	if err != nil {
		return
	}

	output, err = s.royce.GetSingleUser(ctx, newUserId)
	return
}

func (s *service) GetUser(ctx context.Context, input int) (output models.SingleUserData, err error) {
	output, err = s.royce.GetSingleUser(ctx, input)
	return
}

func (s *service) DeleteUser(ctx context.Context, id int) (err error) {
	err = s.royce.DeleteUser(ctx, id)
	return
}

func (s *service) UpdateUser(ctx context.Context, input *models.UpdateUserData) (response models.SingleUserData, err error) {
	err = s.royce.UpdateUser(ctx, input)
	if err != nil {
		return
	}

	response, err = s.royce.GetSingleUser(ctx, input.ID)
	return
}

func (s *service) GetAllUser(ctx context.Context) (response models.AllUsersData, err error) {
	response, err = s.royce.GetAllUser(ctx)

	return
}

// NewService ...
func NewService(royce royce) Service {
	return &service{
		royce: royce,
	}
}
