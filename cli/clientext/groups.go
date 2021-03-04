package clientext

import (
	"context"

	"github.com/GracepointMinistries/hub/client"
	"github.com/antihax/optional"
)

// PageGroups pages through the groups using the given filter
func PageGroups(c *client.APIClient, filter string, handler func(bool, []client.Zgroup) (*client.Zgroup, error)) (*client.Zgroup, error) {
	payload, _, err := c.AdminApi.Zgroups(context.Background(), &client.AdminApiZgroupsOpts{
		Filter: optional.NewString(filter),
	})
	if err != nil {
		return nil, err
	}
	for {
		if payload.Pagination.Cursor == -1 {
			return nil, nil
		}
		found, err := handler(int(payload.Pagination.Limit) > len(payload.Zgroups), payload.Zgroups)
		if err != nil {
			return nil, err
		}
		if found != nil {
			return found, nil
		}
		payload, _, err = c.AdminApi.Zgroups(context.Background(), &client.AdminApiZgroupsOpts{
			Cursor: optional.NewInt64(payload.Pagination.Cursor),
			Limit:  optional.NewInt64(payload.Pagination.Limit),
			Filter: optional.NewString(payload.Pagination.Filter),
		})
		if err != nil {
			return nil, err
		}
	}
}
