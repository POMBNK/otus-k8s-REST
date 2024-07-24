package user

import (
	"context"
	"github.com/POMBNK/kuberRest/internal/service"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	service service.UserService
}

func (s *Server) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {
	return nil, nil
}

func (s *Server) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) FindUserById(ctx context.Context, request FindUserByIdRequestObject) (FindUserByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func New(service service.UserService) *Server {
	return &Server{service: service}
}

func (s *Server) Register(engine fiber.Router) {
	RegisterHandlers(engine, NewStrictHandler(s, []StrictMiddlewareFunc{}))
}
