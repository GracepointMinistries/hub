package utils

import (
	"github.com/gobuffalo/buffalo"
)

// UserIP returns the ip of the requester
func UserIP(c buffalo.Context) string {
	request := c.Request()
	if ip := request.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	if ip := request.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	return request.RemoteAddr
}
