/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

type Setting struct {
	Id int64 `json:"id,omitempty"`
	Script string `json:"script,omitempty"`
	Sheet string `json:"sheet,omitempty"`
	SyncOnce string `json:"syncOnce,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
