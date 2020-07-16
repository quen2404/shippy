// shippy-service-consignment/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro/v2"
	"github.com/quen2404/shippy/shippy-service-consignment/datastore"
	"github.com/quen2404/shippy/shippy-service-consignment/handler"
	pb "github.com/quen2404/shippy/shippy-service-consignment/proto/consignment"
	"github.com/quen2404/shippy/shippy-service-consignment/repository"
	vesselProto "github.com/quen2404/shippy/shippy-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	// Set-up micro instance
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
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
		repository.NewMongoRepository(client.Database("shippy").Collection("consignments")),
		vesselProto.NewVesselService("shippy.service.client", service.Client()),
	)

	// Register handlers
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
