package user

import (
	"context"
	"fmt"
	"github.com/POMBNK/kuberRest/internal/service"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Server struct {
	service service.UserService
	Swagger *openapi3.T
}

func (s *Server) CreateUser(ctx context.Context, request CreateUserRequestObject) (CreateUserResponseObject, error) {

	user := ToCreateUserModel(request.Body)

	userID, err := s.service.CreateUser(ctx, user)
	if err != nil {
		return CreateUser400JSONResponse{
			Code:    http.StatusBadRequest,
			Message: "Error creating user",
		}, err
	}

	return CreateUser201JSONResponse{Id: int64(userID)}, nil
}

func (s *Server) DeleteUser(ctx context.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	err := s.service.DeleteUser(ctx, int(request.UserId))
	if err != nil {
		return DeleteUser400JSONResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Error deleting user with id: %d"),
		}, err
	}

	return DeleteUser200JSONResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("User with id: %d deleted", request.UserId),
	}, nil
}

func (s *Server) FindUserById(ctx context.Context, request FindUserByIdRequestObject) (FindUserByIdResponseObject, error) {
	user, err := s.service.FindUserBydID(ctx, int(request.UserId))
	if err != nil {
		return FindUserById400JSONResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Error finding user with id: %d", request.UserId),
		}, err
	}

	return FindUserById200JSONResponse{
		Email:     &user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Username:  user.Username,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {

	user := ToUpdateUserModel(request.Body)
	err := s.service.UpdateUser(ctx, int(request.UserId), user)

	if err != nil {
		return UpdateUser400JSONResponse{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("Error updating user with id: %d", request.UserId),
		}, err
	}

	return UpdateUser200JSONResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("User with id: %d updated", request.UserId),
	}, nil

}

func New(service service.UserService) *Server {
	swagger, err := GetSwagger()
	if err != nil {
		panic(err)
	}
	return &Server{service: service, Swagger: swagger}
}

func (s *Server) Register(engine fiber.Router) {
	RegisterHandlersWithOptions(engine, NewStrictHandler(s, nil), FiberServerOptions{
		Middlewares: nil,
	})
}
