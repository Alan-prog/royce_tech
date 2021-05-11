package httpclient

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"log"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/pkg/service"
	"my_projects/royce_tech/pkg/service/httpserver"
	"my_projects/royce_tech/tools"
	"net/http"
	"testing"
	"time"
)

const (
	address                  = "127.0.0.1"
	port                     = "8080"
	serverLaunchingWaitSleep = 1 * time.Second

	testAlive            = "testing of the Alive method"
	testCreateUser       = "testing if the CreateUser method"
	testGetSingleUser    = "testing if the GetSingleUser method"
	testDeleteUser       = "testing if the DeleteUser method"
	serviceAlive         = "Alive"
	serviceCreateUser    = "CreateUser"
	serviceGetSingleUser = "GetUser"
	serviceDeleteUser    = "DeleteUser"
)

var (
	nilError tools.ErrorMessage
)

//==============================
//TESTING Alive METHOD
//==============================
func TestClient_Alive(t *testing.T) {
	response := generateValidAliveResponse()
	t.Run(testAlive, func(t *testing.T) {
		serviceMock := new(service.MockService)
		serviceMock.On(serviceAlive, context.Background()).Return(response, nilError)
		server, client := makeClientAndLaunchServer(address, port, serviceMock)
		defer server.Close()
		time.Sleep(serverLaunchingWaitSleep)
		answer, err := client.Alive(context.Background())
		assert.Equal(t, answer, response, "Valid testing")
		assert.NoError(t, err, "unexpected error:", err)
	})
}

//==============================
//TESTING CreateUser METHOD
//==============================
func TestClient_CreateUser(t *testing.T) {
	request := generateValidCreateUserRequest()
	response := generateValidCreateUserData()
	t.Run(testCreateUser, func(t *testing.T) {
		serviceMock := new(service.MockService)
		serviceMock.On(serviceCreateUser, context.Background(), &request).Return(response, nilError)
		server, client := makeClientAndLaunchServer(address, port, serviceMock)
		defer server.Close()
		time.Sleep(serverLaunchingWaitSleep)
		answer, err := client.CreateUser(context.Background(), &request)
		assert.Equal(t, answer, response, "Valid testing")
		assert.NoError(t, err, "unexpected error:", err)
	})

	internalError := tools.NewErrorMessage(errors.New("some error"), "Some human readable", http.StatusInternalServerError)
	t.Run(testCreateUser, func(t *testing.T) {
		serviceMock := new(service.MockService)
		serviceMock.On(serviceCreateUser, context.Background(), &request).Return(response, internalError)
		server, client := makeClientAndLaunchServer(address, port, serviceMock)
		defer server.Close()

		time.Sleep(serverLaunchingWaitSleep)
		_, err := client.CreateUser(context.Background(), &request)
		assert.Equal(t, err, internalError)
	})
}

//==============================
//TESTING GetSingleUser METHOD
//==============================
func TestClient_GetSingleUser(t *testing.T) {
	request := 5
	response := generateValidCreateUserData()
	t.Run(testGetSingleUser, func(t *testing.T) {
		serviceMock := new(service.MockService)
		serviceMock.On(serviceGetSingleUser, context.Background(), request).Return(response, nilError)
		server, client := makeClientAndLaunchServer(address, port, serviceMock)
		defer server.Close()
		time.Sleep(serverLaunchingWaitSleep)
		answer, err := client.GetSingleUser(context.Background(), request)
		assert.Equal(t, answer, response, "Valid testing")
		assert.NoError(t, err, "unexpected error:", err)
	})

	internalError := tools.NewErrorMessage(errors.New("some error"), "Some human readable", http.StatusInternalServerError)
	t.Run(testGetSingleUser, func(t *testing.T) {
		serviceMock := new(service.MockService)
		serviceMock.On(serviceGetSingleUser, context.Background(), request).Return(response, internalError)
		server, client := makeClientAndLaunchServer(address, port, serviceMock)
		defer server.Close()

		time.Sleep(serverLaunchingWaitSleep)
		_, err := client.GetSingleUser(context.Background(), request)
		assert.Equal(t, err, internalError)
	})
}

func generateValidAliveResponse() (response models.AliveResponse) {
	response.Text = "service is okay"
	return
}

func generateValidCreateUserRequest() (response models.CreateUserRequest) {
	name := "Alan"
	description := "golang developer"
	return models.CreateUserRequest{
		Name:        &name,
		Description: &description,
	}
}

func generateValidCreateUserData() (response models.SingleUserData) {
	name := "Alan"
	description := "golang developer"
	return models.SingleUserData{
		Name:        name,
		Description: description,
		CreateAt:    "2009-01-05",
	}
}

func makeClientAndLaunchServer(address, port string, svc service.Service) (server *http.Server, client Service) {
	client = NewPreparedClient(address, port)
	router := httpserver.NewPreparedServer(svc)
	server = &http.Server{
		Handler: router,
		Addr:    address + ":" + port,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("server shut down err: %v", err)
		}
	}()
	return
}
