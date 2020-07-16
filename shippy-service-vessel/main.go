// shippy-service-vessel/main.go
package main

import (
	"github.com/quen2404/shippy/shippy-service-vessel/handler"
	"github.com/quen2404/shippy/shippy-service-vessel/repository"
	"log"

	"github.com/micro/go-micro/v2"
	pb "github.com/quen2404/shippy/shippy-service-vessel/proto/vessel"
)

func main() {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	h := handler.NewHandler(
		repository.NewRepository(vessels),
	)

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
