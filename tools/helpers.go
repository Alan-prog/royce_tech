package tools

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ErrorMessage interface {
	Error() string
	GetCode() int
	GetAll() errorMessage
}

type errorMessage struct {
	Err       string `json:"err"`
	HumanText string `json:"humanText"`
	Code      int    `json:"code"`
}

func (e *errorMessage) Error() string {
	return e.Err
}

func (e *errorMessage) GetCode() int {
	return e.Code
}

func (e *errorMessage) GetAll() errorMessage {
	return *e
}

func EncodeIntoResponseWriter(w http.ResponseWriter, message ErrorMessage) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(message.GetCode())
	json.NewEncoder(w).Encode(message.GetAll())
}

func NewErrorMessage(err error, humanReadable string, code int) ErrorMessage {
	if err == nil {
		return nil
	}
	return &errorMessage{
		Err:       err.Error(),
		HumanText: humanReadable,
		Code:      code,
	}
}

func NewErrorMessageEncodeIntoWriter(err error, humanReadable string, code int) ErrorMessage {
	if err == nil {
		return nil
	}
	return &errorMessage{
		Err:       err.Error(),
		HumanText: humanReadable,
		Code:      code,
	}
}

func DecodeNewErrorMessage(resp *http.Response) ErrorMessage {
	if resp == nil {
		return nil
	}
	var response errorMessage
	if resp.StatusCode != 200 {
		err := json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			return &errorMessage{
				Err: buf.String(),
			}
		}
		return &response
	}
	return nil
}
