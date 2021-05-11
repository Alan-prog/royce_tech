package main

import (
	"context"
	"log"
	"my_projects/royce_tech/pkg/models"
	"my_projects/royce_tech/pkg/royce_tech_app"
	"my_projects/royce_tech/pkg/service"
	"my_projects/royce_tech/pkg/service/httpserver"
	"my_projects/royce_tech/tools/db"
	"net/http"
)

const (
	serverPort = "8080"

	login  = "postgres"
	pass   = "somepass"
	name   = "postgres"
	host   = "127.0.0.1"
	dbPort = uint16(5432)
)

func main() {
	ctx := context.Background()

	dbAdp, err := db.NewDbConnector(ctx, login, pass, host, name, dbPort)
	if err != nil {
		log.Fatalf("error while connecting to db: %v", err)
	}
	defer dbAdp.Close()

	royce := royce_tech_app.NewRoyce(dbAdp)
	svc := service.NewService(royce)
	router := httpserver.NewPreparedServer(svc)
	http.Handle("/", router)

	log.Printf("server starting on port: %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
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
