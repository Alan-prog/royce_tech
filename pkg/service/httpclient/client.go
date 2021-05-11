package httpclient

import (
	"bytes"
	"context"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
	"net/http"
	"strconv"
	"time"
)

const (
	clientTimeout = 15
)

// Service implements Service interface
type Service interface {
	Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err tools.ErrorMessage)
	GetSingleUser(ctx context.Context, input int) (output models.SingleUserData, err tools.ErrorMessage)
}

type client struct {
	cli                    http.Client
	transportAlive         AliveClientTransport
	transportCreateUser    CreateUserClientTransport
	transportGetSingleUser GetSingleUserClientTransport
}

// Alive ...
func (s *client) Alive(ctx context.Context) (output models.AliveResponse, err tools.ErrorMessage) {
	if err = s.transportAlive.EncodeRequest(ctx); err != nil {
		return
	}

	resp, er := s.cli.Get(s.transportAlive.GetPath())
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while making request in AliveClient",
			http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	output, err = s.transportAlive.DecodeResponse(ctx, resp)
	return
}

// CreateUser ...
func (s *client) CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err tools.ErrorMessage) {
	encoded, err := s.transportCreateUser.EncodeRequest(ctx, input)
	if err != nil {
		return
	}

	resp, er := s.cli.Post(s.transportCreateUser.GetPath(), "application/json", bytes.NewBuffer(encoded))
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while making request in CreateUser", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	output, err = s.transportCreateUser.DecodeResponse(ctx, resp)
	return
}

// GetSingleUser...
func (s *client) GetSingleUser(ctx context.Context, input int) (output models.SingleUserData, err tools.ErrorMessage) {
	err = s.transportGetSingleUser.EncodeRequest(ctx)
	if err != nil {
		return
	}

	resp, er := s.cli.Get(s.transportGetSingleUser.GetPath() + "?id=" + strconv.Itoa(input))
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while making request in GetSingleUser", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	output, err = s.transportGetSingleUser.DecodeResponse(ctx, resp)
	return
}

// NewClient the client creator
func NewClient(
	cli http.Client,
	transportAlive AliveClientTransport,
	transportCreateUser CreateUserClientTransport,
	transportGetSingleUser GetSingleUserClientTransport,
) Service {
	return &client{
		cli:                    cli,
		transportCreateUser:    transportCreateUser,
		transportAlive:         transportAlive,
		transportGetSingleUser: transportGetSingleUser,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(address string, port string) Service {
	transportGetUser := NewAliveTransport(HTTP + address + ":" + port + URIPathAlive)
	transportCreateUser := NewCreateUserClientTransport(HTTP + address + ":" + port + URIPathCreateUser)
	transportGetSingleUser := NewGetSingleUserClientTransport(HTTP + address + ":" + port + URIPathUserCommon)
	return NewClient(
		http.Client{
			Timeout: clientTimeout * time.Second,
		},
		transportGetUser,
		transportCreateUser,
		transportGetSingleUser,
	)
}
