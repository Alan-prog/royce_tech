package httpserver

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
	"net/http"
)

// AliveTransport ...
//================================================
// AliveTransport
//================================================
type AliveTransport interface {
	DecodeRequest(ctx context.Context, r *http.Request) (err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.AliveResponse) (err error)
}

type aliveTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *aliveTransport) DecodeRequest(ctx context.Context, r *http.Request) (err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *aliveTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.AliveResponse) (err error) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, err := json.Marshal(response)
	if err != nil {
		err = errors.Wrap(err, "error while marshal Alive response")
		return
	}

	_, err = w.Write(byteResp)
	if err != nil {
		tools.HandleHttpError(w, err)
	}
	return
}

// NewAliveTransport the transport creator for http requests
func NewAliveTransport() AliveTransport {
	return &aliveTransport{}
}

// CreateUserTransport ...
//================================================
// CreateUserTransport
//================================================
type CreateUserTransport interface {
	DecodeRequest(ctx context.Context, r *http.Request) (response models.CreateUserRequest, err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err error)
}

type createUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *createUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (response models.CreateUserRequest,err error) {
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil{
		err = errors.Wrap(err, "error while unmarshal CreateUser request")
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *createUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err error) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, err := json.Marshal(response)
	if err != nil {
		err = errors.Wrap(err, "error while marshal CreateUser response")
		return
	}

	_, err = w.Write(byteResp)
	if err != nil {
		err = errors.Wrap(err, "error while writing response to response writer in CreateUser")
	}
	return
}

// NewCreateUserTransport the transport creator for http requests
func NewCreateUserTransport() CreateUserTransport {
	return &createUserTransport{}
}