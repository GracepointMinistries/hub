/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

// UserWithGroup is a User model with eagerly loaded group data
type UserWithGroup struct {
	Blocked bool `json:"blocked,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Email string `json:"email,omitempty"`
	Group *Group `json:"group,omitempty"`
	Id int64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
