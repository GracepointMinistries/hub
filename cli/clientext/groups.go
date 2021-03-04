package clientext

import (
	"context"

	"github.com/GracepointMinistries/hub/client"
	"github.com/antihax/optional"
)

// PageGroups pages through the groups using the given filter
func PageGroups(c *client.APIClient, filter string, handler func(bool, []client.Group) (*client.Group, error)) (*client.Group, error) {
	payload, _, err := c.AdminApi.Groups(context.Background(), &client.AdminApiGroupsOpts{
		Filter: optional.NewString(filter),
	})
	if err != nil {
		return nil, err
	}
	for {
		if payload.Pagination.Cursor == -1 {
			return nil, nil
		}
		found, err := handler(int(payload.Pagination.Limit) > len(payload.Groups), payload.Groups)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}
		payload, _, err = c.AdminApi.Groups(context.Background(), &client.AdminApiGroupsOpts{
			Cursor: optional.NewInt64(payload.Pagination.Cursor),
			Limit:  optional.NewInt64(payload.Pagination.Limit),
			Filter: optional.NewString(payload.Pagination.Filter),
		})
		if err != nil {
			return nil, err
		}
	}
}
