package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/tools"
	"net/http"
)

type service interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest) (output models.SingleUserData, err error)
	GetUser(ctx context.Context, input int) (output models.SingleUserData, err error)
	DeleteUser(ctx context.Context, input int) (err error)
	UpdateUser(ctx context.Context, input *models.UpdateUserData) (output models.SingleUserData, err error)
	GetAllUser(ctx context.Context) (output models.AllUsersData, err error)
}

//================================================
// AliveServer
//================================================
type aliveServer struct {
	transport AliveTransport
	service   service
}

// ServeHTTP implements http.Handler.
func (s *aliveServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.transport.DecodeRequest(r.Context(), r)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	response, err := s.service.Alive()
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, &response); err != nil {
		tools.HandleHttpError(w, err)
		return
	}
}

// NewAliveServer the server creator
func NewAliveServer(transport AliveTransport, service service) http.HandlerFunc {
	ls := aliveServer{
		transport: transport,
		service:   service,
	}
	return ls.ServeHTTP
}

//================================================
// CreateUser
//================================================
type createUserServer struct {
	transport CreateUserTransport
	service   service
}

// ServeHTTP implements http.Handler.
func (s *createUserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := s.transport.DecodeRequest(r.Context(), r)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	response, err := s.service.CreateUser(r.Context(), &req)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, &response); err != nil {
		tools.HandleHttpError(w, err)
		return
	}
}

// NewCreateUserServer the server creator
func NewCreateUserServer(transport CreateUserTransport, service service) http.HandlerFunc {
	ls := createUserServer{
		transport: transport,
		service:   service,
	}
	return ls.ServeHTTP
}

//================================================
// GetUser
//================================================
type getUserServer struct {
	transport GetUserTransport
	service   service
}

// ServeHTTP implements http.Handler.
func (s *getUserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := s.transport.DecodeRequest(r.Context(), r)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	response, err := s.service.GetUser(r.Context(), id)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, &response); err != nil {
		tools.HandleHttpError(w, err)
		return
	}
}

// NewGetUserServer the server creator
func NewGetUserServer(transport GetUserTransport, service service) http.HandlerFunc {
	ls := getUserServer{
		transport: transport,
		service:   service,
	}
	return ls.ServeHTTP
}

//================================================
// DeleteUser
//================================================
type deleteUserServer struct {
	transport DeleteUserTransport
	service   service
}

// ServeHTTP implements http.Handler.
func (s *deleteUserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := s.transport.DecodeRequest(r.Context(), r)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	err = s.service.DeleteUser(r.Context(), id)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w); err != nil {
		tools.HandleHttpError(w, err)
		return
	}
}

// NewDeleteUserServer the server creator
func NewDeleteUserServer(transport DeleteUserTransport, service service) http.HandlerFunc {
	ls := deleteUserServer{
		transport: transport,
		service:   service,
	}
	return ls.ServeHTTP
}

//================================================
// UpdateUser
//================================================
type updateUserServer struct {
	transport UpdateUserTransport
	service   service
}

// ServeHTTP implements http.Handler.
func (s *updateUserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := s.transport.DecodeRequest(r.Context(), r)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	response, err := s.service.UpdateUser(r.Context(), &req)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, &response); err != nil {
		tools.HandleHttpError(w, err)
		return
	}
}

// NewUpdateUserServer the server creator
func NewUpdateUserServer(transport UpdateUserTransport, service service) http.HandlerFunc {
	ls := updateUserServer{
		transport: transport,
		service:   service,
	}
	return ls.ServeHTTP
}

//================================================
// GetAllUserServer
//================================================
type getAllUserServer struct {
	transport GetAllUserTransport
	service   service
}

// ServeHTTP implements http.Handler.
func (s *getAllUserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := s.transport.DecodeRequest(r.Context(), r)
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	response, err := s.service.GetAllUser(r.Context())
	if err != nil {
		tools.HandleHttpError(w, err)
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, response); err != nil {
		tools.HandleHttpError(w, err)
		return
	}
}

// NewGetAllUserServer the server creator
func NewGetAllUserServer(transport GetAllUserTransport, service service) http.HandlerFunc {
	ls := getAllUserServer{
		transport: transport,
		service:   service,
	}
	return ls.ServeHTTP
}

// NewPreparedServer ...
func NewPreparedServer(svc service) *mux.Router {
	aliveTransport := NewAliveTransport()
	createUserTransport := NewCreateUserTransport()
	getUserTransport := NewGetUserTransport()
	deleteUserTransport := NewDeleteUserTransport()
	updateUserTransport := NewUpdateUserTransport()
	getAllUserTransport := NewGetAllUsersTransport()

	return MakeRouter(
		[]*HandlerSettings{
			{
				Path:    URIPathGetAlive,
				Method:  http.MethodGet,
				Handler: NewAliveServer(aliveTransport, svc),
			},
			{
				Path:    URIPathNewUser,
				Method:  http.MethodPost,
				Handler: NewCreateUserServer(createUserTransport, svc),
			},
			{
				Path:    URIPathGetUser,
				Method:  http.MethodGet,
				Handler: NewGetUserServer(getUserTransport, svc),
			},
			{
				Path:    URIPathGetUser,
				Method:  http.MethodDelete,
				Handler: NewDeleteUserServer(deleteUserTransport, svc),
			},
			{
				Path:    URIPathGetUser,
				Method:  http.MethodPut,
				Handler: NewUpdateUserServer(updateUserTransport, svc),
			},
			{
				Path:    URIPathGetAllUser,
				Method:  http.MethodGet,
				Handler: NewGetAllUserServer(getAllUserTransport, svc),
			},
		},
	)
}
