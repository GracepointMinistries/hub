/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// AdminZgroupResponsePayload contains the queried zgroup and its users
type AdminZgroupResponsePayload struct {
	Users []User `json:"users,omitempty"`
	Zgroup *Zgroup `json:"zgroup,omitempty"`
}
