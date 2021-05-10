package service

import (
	"context"
	"my_projects/royce_tech/pkg/models"
)

type royce interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest)(output int32, err error)
	GetSingleUser(ctx context.Context, id int32)(output models.SingleUserData, err error)
}

type Service interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest)(output models.SingleUserData, err error)
}

type service struct {
	royce royce
}

func (s *service) Alive() (output models.AliveResponse, err error) {
	output, err = s.royce.Alive()
	return
}

func (s *service) CreateUser(ctx context.Context, input *models.CreateUserRequest)(output models.SingleUserData, err error){
	newUserId, err := s.royce.CreateUser(ctx, input)
	if err != nil{
		return
	}

	output, err = s.royce.GetSingleUser(ctx, newUserId)
	return
}

// NewService ...
func NewService(royce royce) Service {
	return &service{
		royce: royce,
	}
}
