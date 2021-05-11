package httpclient

import (
	"context"
	"encoding/json"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
	"net/http"
)

//================================================
// AliveClientTransport
//================================================
type AliveClientTransport interface {
	EncodeRequest(ctx context.Context) (err tools.ErrorMessage)
	DecodeResponse(ctx context.Context, r *http.Response) (response models.AliveResponse, err tools.ErrorMessage)
	GetPath() (response string)
}

type aliveClientTransport struct {
	path string
}

// EncodeRequest method for encoding requests on client side
func (t *aliveClientTransport) EncodeRequest(ctx context.Context) (err tools.ErrorMessage) {
	return
}

// DecodeResponse method for decoding response on client side
func (t *aliveClientTransport) DecodeResponse(ctx context.Context, r *http.Response) (response models.AliveResponse, err tools.ErrorMessage) {
	if r.StatusCode != http.StatusOK {
		err = tools.DecodeNewErrorMessage(r)
		return
	}

	er := json.NewDecoder(r.Body).Decode(&response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while unmarshal AliveClient response", 0)
		return
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
	EncodeRequest(ctx context.Context, req *models.CreateUserRequest) (response []byte, err tools.ErrorMessage)
	DecodeResponse(ctx context.Context, r *http.Response) (response models.SingleUserData, err tools.ErrorMessage)
	GetPath() (response string)
}

type createUserClientTransport struct {
	path string
}

// EncodeRequest method for encoding requests on client side
func (t *createUserClientTransport) EncodeRequest(ctx context.Context, req *models.CreateUserRequest) (response []byte, err tools.ErrorMessage) {
	response, er := json.Marshal(req)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while marshal CreateUserClient request", 0)
	}
	return
}

// DecodeResponse method for decoding response on client side
func (t *createUserClientTransport) DecodeResponse(ctx context.Context, r *http.Response) (response models.SingleUserData, err tools.ErrorMessage) {
	if r.StatusCode != http.StatusOK {
		err = tools.DecodeNewErrorMessage(r)
		return
	}

	er := json.NewDecoder(r.Body).Decode(&response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while unmarshal CreateUserClient response", 0)
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

//================================================
// GetSingleUserClientTransport
//================================================
type GetSingleUserClientTransport interface {
	EncodeRequest(ctx context.Context) (err tools.ErrorMessage)
	DecodeResponse(ctx context.Context, r *http.Response) (response models.SingleUserData, err tools.ErrorMessage)
	GetPath() (response string)
}

type getSingleUserClientTransport struct {
	path string
}

// EncodeRequest method for encoding requests on client side
func (t *getSingleUserClientTransport) EncodeRequest(ctx context.Context) (err tools.ErrorMessage) {
	return
}

// DecodeResponse method for decoding response on client side
func (t *getSingleUserClientTransport) DecodeResponse(ctx context.Context, r *http.Response) (response models.SingleUserData, err tools.ErrorMessage) {
	if r.StatusCode != http.StatusOK {
		err = tools.DecodeNewErrorMessage(r)
		return
	}

	er := json.NewDecoder(r.Body).Decode(&response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while unmarshal GetSingleUserClient response", 0)
	}
	return
}

func (t *getSingleUserClientTransport) GetPath() string {
	return t.path
}

// NewGetSingleUserClientTransport the transport creator for http requests
func NewGetSingleUserClientTransport(
	pathTemplate string,
) GetSingleUserClientTransport {
	return &getSingleUserClientTransport{
		path: pathTemplate,
	}
}
