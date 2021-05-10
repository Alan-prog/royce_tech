package service

import "my_projects/royce_tech/pkg/models"

type royce interface {
	Alive() (output models.AliveResponse, err error)
}

type Service interface {
	Alive() (output models.AliveResponse, err error)
}

type service struct {
	royce royce
}

func (s *service) Alive() (output models.AliveResponse, err error) {
	output, err = s.royce.Alive()
	return
}

// NewService ...
func NewService(royce royce) Service {
	return &service{
		royce: royce,
	}
}
