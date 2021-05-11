package httpserver

import (
	"context"
	"encoding/json"
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
	DecodeRequest(ctx context.Context, r *http.Request) (err tools.ErrorMessage)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.AliveResponse) (err tools.ErrorMessage)
}

type aliveTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *aliveTransport) DecodeRequest(ctx context.Context, r *http.Request) (err tools.ErrorMessage) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *aliveTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.AliveResponse) (err tools.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, er := json.Marshal(response)
	if er != nil {
		err = tools.NewErrorMessage(err, "error while marshal Alive response", http.StatusInternalServerError)
		return
	}

	_, er = w.Write(byteResp)
	if er != nil {
		err = tools.NewErrorMessage(err, "error while marshal Alive response", http.StatusInternalServerError)
		return
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
	DecodeRequest(ctx context.Context, r *http.Request) (response models.CreateUserRequest, err tools.ErrorMessage)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err tools.ErrorMessage)
}

type createUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *createUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (response models.CreateUserRequest, err tools.ErrorMessage) {
	er := json.NewDecoder(r.Body).Decode(&response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while unmarshal CreateUser request", http.StatusInternalServerError)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *createUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err tools.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, er := json.Marshal(response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while marshal CreateUser response", http.StatusInternalServerError)
		return
	}

	_, er = w.Write(byteResp)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while writing response to response writer in CreateUser",
			http.StatusInternalServerError)
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
	DecodeRequest(ctx context.Context, r *http.Request) (id int, err tools.ErrorMessage)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err tools.ErrorMessage)
}

type getUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *getUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (id int, err tools.ErrorMessage) {
	id, er := strconv.Atoi(r.URL.Query().Get("id"))
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while parsing query in GetUser", http.StatusInternalServerError)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err tools.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, er := json.Marshal(response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while marshal GetUser response",
			http.StatusInternalServerError)
		return
	}

	_, er = w.Write(byteResp)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while writing response to response writer in GetUser",
			http.StatusInternalServerError)
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
	DecodeRequest(ctx context.Context, r *http.Request) (id int, err tools.ErrorMessage)
	EncodeResponse(ctx context.Context, w http.ResponseWriter) (err tools.ErrorMessage)
}

type deleteUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *deleteUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (id int, err tools.ErrorMessage) {
	id, er := strconv.Atoi(r.URL.Query().Get("id"))
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while parsing query in DeleteUser",
			http.StatusInternalServerError)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *deleteUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter) (err tools.ErrorMessage) {
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
	DecodeRequest(ctx context.Context, r *http.Request) (response models.UpdateUserData, err tools.ErrorMessage)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err tools.ErrorMessage)
}

type updateUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *updateUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (response models.UpdateUserData, err tools.ErrorMessage) {
	er := json.NewDecoder(r.Body).Decode(&response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while unmarshal CreateUser request",
			http.StatusInternalServerError)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *updateUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response *models.SingleUserData) (err tools.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, er := json.Marshal(response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while marshal UpdateUser response",
			http.StatusInternalServerError)
		return
	}

	_, er = w.Write(byteResp)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while writing response to response writer in UpdateUser",
			http.StatusInternalServerError)
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
	DecodeRequest(ctx context.Context, r *http.Request) (err tools.ErrorMessage)
	EncodeResponse(ctx context.Context, w http.ResponseWriter, response models.AllUsersData) (err tools.ErrorMessage)
}

type getAllUserTransport struct {
}

// DecodeRequest method for decoding requests on server side
func (t *getAllUserTransport) DecodeRequest(ctx context.Context, r *http.Request) (err tools.ErrorMessage) {
	return
}

// EncodeResponse method for encoding response on server side
func (t *getAllUserTransport) EncodeResponse(ctx context.Context, w http.ResponseWriter, response models.AllUsersData) (err tools.ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	byteResp, er := json.Marshal(response)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while marshal GetAllUser response",
			http.StatusInternalServerError)
		return
	}

	_, er = w.Write(byteResp)
	if er != nil {
		err = tools.NewErrorMessage(er, "Error while writing response to response writer in GetAllUser",
			http.StatusInternalServerError)
	}
	return
}

// NewGetAllUsersTransport the transport creator for http requests
func NewGetAllUsersTransport() GetAllUserTransport {
	return &getAllUserTransport{}
}
