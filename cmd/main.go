package main

import (
	"context"
	"log"
	"net/http"

	"my_projects/royce_tech/pkg/royce_tech_app"
	"my_projects/royce_tech/pkg/service"
	"my_projects/royce_tech/pkg/service/httpserver"
	"my_projects/royce_tech/tools/db"
)

const (
	serverPort = "8080"

	login = "postgres"
	pass = "somepass"
	name = "postgres"
	host = "127.0.0.1"
	dbPort = uint16(6080)
)

const(
	createTableRequest = `create table human_resources
(
	id serial not null,
	name varchar(256),
	dob date,
	address varchar(256),
	description varchar(512),
	created_at date not null,
	updated_at date
);

create unique index human_resources_id_uindex
	on human_resources (id);

alter table human_resources
	add constraint table_name_pk
		primary key (id);`
)

func main() {
	ctx := context.Background()

	dbAdp, err := db.NewDbConnector(ctx, login, pass, host,name,dbPort)
	if err != nil{
		log.Fatalf("error while connecting to db: %v",err)
	}

	royce := royce_tech_app.NewRoyce(dbAdp)
	svc := service.NewService(royce)
	router := httpserver.NewPreparedServer(svc)

	http.Handle("/", router)

	log.Printf("server starting on port: %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+ serverPort, nil))
}
