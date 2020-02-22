package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/ennc0d3/madrush/pkg/swagger/server/restapi"
	"github.com/ennc0d3/madrush/pkg/swagger/server/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func serve() {
	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewMadrushAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()
	server.Port = 8307

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)

	// Start listening using having the handlers and port
	// already set up.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

//Health route returns OK
func Health(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}

//GetHelloUser returns Hello + your name
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}

func main() {
	serve()
}
