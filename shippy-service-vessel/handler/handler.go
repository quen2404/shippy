package handler

import (
	pb "github.com/quen2404/shippy/shippy-service-vessel/proto/vessel"
	"github.com/quen2404/shippy/shippy-service-vessel/repository"
	"golang.org/x/net/context"
)

// Our grpc service handler
type Handler struct {
	repository repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo}
}

func (s *Handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repository.FindAvailable(ctx, repository.MarshalSpecification(req))
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = repository.UnmarshalVessel(vessel)
	return nil
}

// Create a new vessel
func (s *Handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(ctx, repository.MarshalVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
