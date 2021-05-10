package httpserver

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
	"net/http"
	"strconv"
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
func (t *createUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (response models.CreateUserRequest, err error) {
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
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

// GetUserTransport ...
//================================================
// GetUserTransport
//================================================
type GetUserTransport interface {
	DecodeRequest(ctx context.Context, r *http.Request) (id int, err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err error)
}

type getUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *getUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		err = errors.Wrap(err, "error while parsing query in GetUser")
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err error) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, err := json.Marshal(response)
	if err != nil {
		err = errors.Wrap(err, "error while marshal GetUser response")
		return
	}

	_, err = w.Write(byteResp)
	if err != nil {
		err = errors.Wrap(err, "error while writing response to response writer in GetUser")
	}
	return
}

// NewGetUserTransport the transport creator for http requests
func NewGetUserTransport() GetUserTransport {
	return &getUserTransport{}
}

// DeleteUserTransport ...
//================================================
// DeleteUserTransport
//================================================
type DeleteUserTransport interface {
	DecodeRequest(ctx context.Context, r *http.Request) (id int, err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter) (err error)
}

type deleteUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *deleteUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (id int, err error) {
	id, err = strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		err = errors.Wrap(err, "error while parsing query in DeleteUser")
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *deleteUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter) (err error) {
	return
}

// NewDeleteUserTransport the transport creator for http requests
func NewDeleteUserTransport() DeleteUserTransport {
	return &deleteUserTransport{}
}

// UpdateUserTransport ...
//================================================
// UpdateUserTransport
//================================================
type UpdateUserTransport interface {
	DecodeRequest(ctx context.Context, r *http.Request) (response models.UpdateUserData, err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err error)
}

type updateUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *updateUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (response models.UpdateUserData, err error) {
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		err = errors.Wrap(err, "error while unmarshal CreateUser request")
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *updateUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err error) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, err := json.Marshal(response)
	if err != nil {
		err = errors.Wrap(err, "error while marshal UpdateUser response")
		return
	}

	_, err = w.Write(byteResp)
	if err != nil {
		err = errors.Wrap(err, "error while writing response to response writer in UpdateUser")
	}
	return
}

// NewUpdateUserTransport the transport creator for http requests
func NewUpdateUserTransport() UpdateUserTransport {
	return &updateUserTransport{}
}

// GetAllUserTransport ...
//================================================
// GetAllUserTransport
//================================================
type GetAllUserTransport interface {
	DecodeRequest(ctx context.Context, r *http.Request) (err error)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response models.AllUsersData) (err error)
}

type getAllUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *getAllUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (err error) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getAllUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response models.AllUsersData) (err error) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, err := json.Marshal(response)
	if err != nil {
		err = errors.Wrap(err, "error while marshal GetAllUser response")
		return
	}

	_, err = w.Write(byteResp)
	if err != nil {
		err = errors.Wrap(err, "error while writing response to response writer in GetAllUser")
	}
	return
}

// NewGetAllUsersTransport the transport creator for http requests
func NewGetAllUsersTransport() GetAllUserTransport {
	return &getAllUserTransport{}
}
