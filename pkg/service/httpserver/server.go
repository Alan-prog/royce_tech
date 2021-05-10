package httpserver

import (
	"context"
	"github.com/gorilla/mux"
	"my_projects/royce_tech/pkg/models"
	"net/http"
)

type service interface {
	Alive() (output models.AliveResponse, err error)
	CreateUser(ctx context.Context, input *models.CreateUserRequest)(output models.SingleUserData, err error)
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
		return
	}

	response, err := s.service.Alive()
	if err != nil {
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, &response); err != nil {
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
		return
	}

	response, err := s.service.CreateUser(r.Context(), &req)
	if err != nil {
		return
	}

	if err := s.transport.EncodeResponse(r.Context(), w, &response); err != nil {
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

// NewPreparedServer ...
func NewPreparedServer(svc service) *mux.Router {
	aliveTransport := NewAliveTransport()
	createUserTransport := NewCreateUserTransport()

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
		},
	)
}
