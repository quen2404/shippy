// shippy-service-vessel/main.go
package main

import (
	"context"
	"github.com/quen2404/shippy/shippy-service-vessel/datastore"
	"github.com/quen2404/shippy/shippy-service-vessel/handler"
	"github.com/quen2404/shippy/shippy-service-vessel/repository"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	pb "github.com/quen2404/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := datastore.CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	h := handler.NewHandler(
		repository.NewRepository(client.Database("shippy").Collection("vessels")),
	)

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
