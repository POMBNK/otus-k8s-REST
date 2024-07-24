package entity

import openapi_types "github.com/oapi-codegen/runtime/types"

type User struct {
	Email     *openapi_types.Email `json:"email,omitempty"`
	FirstName *string              `json:"firstName,omitempty"`
	Id        *int64               `json:"id,omitempty"`
	LastName  *string              `json:"lastName,omitempty"`
	Phone     *string              `json:"phone,omitempty"`
	Username  *string              `json:"username,omitempty"`
}
