package handler

import (
	pb "github.com/quen2404/shippy/shippy-service-vessel/proto/vessel"
	"github.com/quen2404/shippy/shippy-service-vessel/repository"
	"golang.org/x/net/context"
)

// Our grpc service handler
type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo}
}

func (s *Handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	// Set the vessel as part of the response message type
	res.Vessel = vessel
	return nil
}
