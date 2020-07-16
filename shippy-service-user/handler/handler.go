package handler

import (
	"context"
	"errors"
	"github.com/quen2404/shippy/shippy-service-user/repository"
	"github.com/quen2404/shippy/shippy-service-user/services"

	pb "github.com/quen2404/shippy/shippy-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authable interface {
	Decode(token string) (*services.CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type Handler struct {
	repository   repository.Repository
	tokenService authable
}

func NewHandler(repository repository.Repository, tokenService authable) *Handler {
	return &Handler{
		repository:   repository,
		tokenService: tokenService,
	}
}

func (s *Handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	result, err := s.repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}

	user := repository.UnmarshalUser(result)
	res.User = user

	return nil
}

func (s *Handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}

	users := repository.UnmarshalUserCollection(results)
	res.Users = users

	return nil
}

func (s *Handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := s.tokenService.Encode(req)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (s *Handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := s.repository.Create(ctx, repository.MarshalUser(req)); err != nil {
		return err
	}

	// Strip the password back out, so's we're not returning it
	req.Password = ""
	res.User = req

	return nil
}

func (s *Handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true
	return nil
}
