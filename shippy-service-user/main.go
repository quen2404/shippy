package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/quen2404/shippy/shippy-service-user/database"
	"github.com/quen2404/shippy/shippy-service-user/handler"
	pb "github.com/quen2404/shippy/shippy-service-user/proto/user"
	"github.com/quen2404/shippy/shippy-service-user/repository"
	"github.com/quen2404/shippy/shippy-service-user/services"
	"log"
)

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := database.CreateConnection()
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	repo := repository.NewPostgresRepository(db)

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	h := handler.NewHandler(repo, services.NewTokenService(repo))

	// Register handler
	if err := pb.RegisterUserServiceHandler(srv.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
