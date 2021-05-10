package httpserver

import (
	"context"
	"encoding/json"
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
		tools.HandleHttpError(w, err)
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
