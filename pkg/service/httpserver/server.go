package httpserver

import (
	"github.com/gorilla/mux"
	"my_projects/royce_tech/pkg/models"
	"net/http"
)

type service interface {
	Alive() (output models.AliveResponse, err error)
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

// NewPreparedServer ...
func NewPreparedServer(svc service) *mux.Router {
	aliveTransport := NewAliveTransport()

	return MakeRouter(
		[]*HandlerSettings{
			{
				Path:    URIPathGetAlive,
				Method:  http.MethodGet,
				Handler: NewAliveServer(aliveTransport, svc),
			},
		},
	)
}
