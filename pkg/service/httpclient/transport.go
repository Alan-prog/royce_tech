package httpclient

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"my_projects/royce_tech/pkg/models"
	"net/http"
)

//================================================
// AliveClientTransport
//================================================
type AliveClientTransport interface {
	EncodeRequest(ctx context.Context) (err error)
	DecodeResponse(ctx context.Context, r *http.Response) (response models.AliveResponse, err error)
	GetPath() (response string)
}

type aliveClientTransport struct {
	path string
}

// EncodeRequest method for encoding requests on client side
func (t *aliveClientTransport) EncodeRequest(ctx context.Context) (err error) {
	return
}

// DecodeResponse method for decoding response on client side
func (t *aliveClientTransport) DecodeResponse(ctx context.Context, r *http.Response) (response models.AliveResponse, err error) {
	if r.StatusCode != http.StatusOK {
		err = errors.Wrap(err, "bad status code from server")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		err = errors.Wrap(err, "error while unmarshal AliveClient response")
	}
	return
}

func (t *aliveClientTransport) GetPath() string {
	return t.path
}

// NewAliveTransport the transport creator for http requests
func NewAliveTransport(
	pathTemplate string,
) AliveClientTransport {
	return &aliveClientTransport{
		path: pathTemplate,
	}
}

//================================================
// CreateUserClientTransport
//================================================
type CreateUserClientTransport interface {
	EncodeRequest(ctx context.Context, req *models.CreateUserRequest) (response []byte, err error)
	DecodeResponse(ctx context.Context, r *http.Response) (response models.SingleUserData, err error)
	GetPath() (response string)
}

type createUserClientTransport struct {
	path string
}

// EncodeRequest method for encoding requests on client side
func (t *createUserClientTransport) EncodeRequest(ctx context.Context, req *models.CreateUserRequest) (response []byte, err error) {
	response, err = json.Marshal(req)
	if err != nil {
		err = errors.Wrap(err, "error while marshal CreateUserClient request")
	}
	return
}

// DecodeResponse method for decoding response on client side
func (t *createUserClientTransport) DecodeResponse(ctx context.Context, r *http.Response) (response models.SingleUserData, err error) {
	if r.StatusCode != http.StatusOK {
		err = errors.Wrap(err, "bad status code from server")
		return
	}
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		err = errors.Wrap(err, "error while unmarshal CreateUserClient response")
	}
	return
}

func (t *createUserClientTransport) GetPath() string {
	return t.path
}

// NewCreateUserClientTransport the transport creator for http requests
func NewCreateUserClientTransport(
	pathTemplate string,
) CreateUserClientTransport {
	return &createUserClientTransport{
		path: pathTemplate,
	}
}
