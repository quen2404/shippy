package services

import (
	pb "github.com/quen2404/shippy/shippy-service-user/proto/user"
	"github.com/quen2404/shippy/shippy-service-user/repository"
)

type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type CustomClaims struct {
	User pb.User
}

type TokenService struct {
	repo repository.Repository
}

func NewTokenService(repo repository.Repository) *TokenService {
	return &TokenService{repo}
}

func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	return &CustomClaims{}, nil
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	return "", nil
}
