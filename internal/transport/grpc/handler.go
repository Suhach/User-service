package grpc

import (
	userpb "github.com/Suhach/projekt-protoc/proto/user"
	"github.com/your-org/users-service/internal/user"
)

type Handler struct {
	svc *user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc *user.Service) *Handler {
	return &Handler{svc: svc}
}
