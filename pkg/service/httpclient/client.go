package httpclient

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"my_projects/royce_tech/pkg/models"
	"net/http"
	"time"
)

const (
	clientTimeout = 15
)

// Service implements Service interface
type Service interface {
	Alive(ctx context.Context) (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err error)
}

type client struct {
	cli                 http.Client
	transportAlive      AliveClientTransport
	transportCreateUser CreateUserClientTransport
}

// Alive ...
func (s *client) Alive(ctx context.Context) (output models.AliveResponse, err error) {
	if err = s.transportAlive.EncodeRequest(ctx); err != nil {
		return
	}

	resp, err := s.cli.Get(s.transportAlive.GetPath())
	if err != nil {
		err = errors.Wrap(err, "error while making request in AliveClient")
		return
	}
	defer resp.Body.Close()

	output, err = s.transportAlive.DecodeResponse(ctx, resp)

	return
}

// CreateUser ...
func (s *client) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err error) {
	encoded, err := s.transportCreateUser.EncodeRequest(ctx, input)
	if err != nil {
		return
	}

	resp, err := s.cli.Post(s.transportCreateUser.GetPath(), "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		err = errors.Wrap(err, "error while making request in CreateUser")
		return
	}
	defer resp.Body.Close()

	output, err = s.transportCreateUser.DecodeResponse(ctx, resp)
	return
}

// NewClient the client creator
func NewClient(
	cli http.Client,
	transportAlive AliveClientTransport,
	transportCreateUser CreateUserClientTransport,
) Service {
	return &client{
		cli:                 cli,
		transportCreateUser: transportCreateUser,
		transportAlive:      transportAlive,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(address string, port string) Service {
	transportGetUser := NewAliveTransport(HTTP + address + ":" + port + URIPathAlive)
	transportCreateUser := NewCreateUserClientTransport(HTTP + address + ":" + port + URIPathCreateUser)
	return NewClient(
		http.Client{
			Timeout: clientTimeout * time.Second,
		},
		transportGetUser,
		transportCreateUser,
	)
}
