package httpclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/pkg/service"
	"my_projects/royce_tech/pkg/service/httpserver"
	"net/http"
	"testing"
	"time"
)

const (
	address                  = "127.0.0.1"
	port                     = "8080"
	serverLaunchingWaitSleep = 1 * time.Second

	testAlive         = "testing of the Alive method"
	testCreateUser    = "testing if the CreateUser method"
	serviceAlive      = "Alive"
	serviceCreateUser = "CreateUser"
)

var (
	nilError error
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
	response := generateValidCreateUserResponse()
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

	//internalError := error.NewError(500,"internal error")
	//request = generateInvalidRequest()
	//response = generateInvalidResponse()
	//t.Run(testGetOrders, func(t *testing.T) {
	//	serviceMock := new(dataService.MockService)
	//	serviceMock.On(serviceGetOrders, context.Background(), &request).
	//		Return(response,internalError)
	//	server, client := makeServerClient(cfgPort, serviceMock)
	//	defer func() {
	//		err := server.Shutdown()
	//		if err != nil {
	//			log.Printf("server shut down err: %v", err)
	//		}
	//	}()
	//	time.Sleep(serverLaunchingWaitSleep)
	//	_, err := client.GetOrders(context.Background(), &request)
	//	assert.Equal(t, err, internalError)
	//})
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

func generateValidCreateUserResponse() (response models.SingleUserData) {
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
