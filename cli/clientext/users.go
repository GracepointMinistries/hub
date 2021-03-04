package clientext

import (
	"context"

	"github.com/GracepointMinistries/hub/client"
	"github.com/antihax/optional"
)

// PageUsers pages through the users using the given filter
func PageUsers(c *client.APIClient, filter string, handler func(bool, []client.UserWithZgroup) (*client.UserWithZgroup, error)) (*client.UserWithZgroup, error) {
	payload, _, err := c.AdminApi.Users(context.Background(), &client.AdminApiUsersOpts{
		Filter: optional.NewString(filter),
	})
	if err != nil {
		return nil, err
	}
	for {
		if payload.Pagination.Cursor == -1 {
			return nil, nil
		}
		found, err := handler(int(payload.Pagination.Limit) > len(payload.Users), payload.Users)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}
		payload, _, err = c.AdminApi.Users(context.Background(), &client.AdminApiUsersOpts{
			Cursor: optional.NewInt64(payload.Pagination.Cursor),
			Limit:  optional.NewInt64(payload.Pagination.Limit),
			Filter: optional.NewString(payload.Pagination.Filter),
		})
		if err != nil {
			return nil, err
		}
	}
}

// AllUsers returns an unpaginated list of all users based on the given filter
func AllUsers(c *client.APIClient, filter string) ([]client.UserWithZgroup, error) {
	payload, _, err := c.AdminApi.Users(context.Background(), &client.AdminApiUsersOpts{
		Filter: optional.NewString(filter),
		Limit:  optional.NewInt64(-1),
	})
	return payload.Users, err
}
