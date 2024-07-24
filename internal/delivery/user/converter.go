package user

import "github.com/POMBNK/kuberRest/internal/entity"

func ToCreateUserModel(request *CreateUserJSONRequestBody) entity.CreateUser {
	return entity.CreateUser{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Username:  request.Username,
	}
}

func ToUpdateUserModel(request *UpdateUserJSONRequestBody) entity.User {
	return entity.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Username:  request.Username,
	}
}
