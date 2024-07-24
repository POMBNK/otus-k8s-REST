package entity

import openapi_types "github.com/oapi-codegen/runtime/types"

type User struct {
	ID        int64
	Email     openapi_types.Email `json:"email,omitempty"`
	FirstName string              `json:"firstName,omitempty"`
	LastName  string              `json:"lastName,omitempty"`
	Phone     string              `json:"phone,omitempty"`
	Username  string              `json:"username,omitempty"`
}

type UpdateUser struct {
	Email     openapi_types.Email `json:"email,omitempty"`
	FirstName string              `json:"firstName,omitempty"`
	LastName  string              `json:"lastName,omitempty"`
	Phone     string              `json:"phone,omitempty"`
	Username  string              `json:"username,omitempty"`
}

type CreateUser struct {
	Email     *openapi_types.Email `json:"email,omitempty"`
	FirstName string               `json:"firstName,omitempty"`
	LastName  string               `json:"lastName,omitempty"`
	Phone     string               `json:"phone,omitempty"`
	Username  string               `json:"username,omitempty"`
}
