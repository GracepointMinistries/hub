/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

type Setting struct {
	Id int64 `json:"id,omitempty"`
	Sheet string `json:"sheet,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
