package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/GracepointMinistries/hub/client"
	"github.com/antihax/optional"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	all bool
)

func pageUsers(c *client.APIClient, filter string, handler func(bool, []client.UserWithZgroup) (*client.UserWithZgroup, error)) (*client.UserWithZgroup, error) {
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

func userOptions(final bool, users []client.UserWithZgroup) (map[int]client.UserWithZgroup, []string) {
	lookup := make(map[int]client.UserWithZgroup, len(users))
	options := make([]string, len(users)+1)
	for i, user := range users {
		name := fmt.Sprintf("%s %s", bold(user.Name), noticef("[ID: %d]", user.Id))
		if user.Blocked {
			name = fmt.Sprintf("%s %s %s", bold(user.Name), warning("(blocked)"), noticef("[ID: %d]", user.Id))
		}
		options[i] = name
		lookup[i] = user
	}
	if final {
		options[len(users)] = exitOption
	} else {
		options[len(users)] = retrieveMoreOption
	}
	return lookup, options
}

func selectUser(c *client.APIClient, filter string) (*client.UserWithZgroup, error) {
	return pageUsers(c, filter, func(final bool, users []client.UserWithZgroup) (*client.UserWithZgroup, error) {
		lookup, options := userOptions(final, users)
		prompt := promptui.Select{
			Label:        bold("Select User"),
			Items:        options,
			HideSelected: true,
			Stdout:       &bellSkipper{},
		}
		index, result, err := prompt.Run()
		if err != nil {
			return nil, err
		}
		if result == exitOption || result == retrieveMoreOption {
			return nil, nil
		}
		if user, found := lookup[index]; found {
			return &user, nil
		}
		return nil, nil
	})
}

func allUsers(c *client.APIClient, filter string) ([]client.UserWithZgroup, error) {
	payload, _, err := c.AdminApi.Users(context.Background(), &client.AdminApiUsersOpts{
		Filter: optional.NewString(filter),
		Limit:  optional.NewInt64(-1),
	})
	return payload.Users, err
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Retrieve information about a user",
	Run: func(cmd *cobra.Command, args []string) {
		c := newClient()
		if all {
			users, err := allUsers(c, filter)
			checkError(err)
			dumpUsersWithZgroup(users...)
			return
		}
		if user != 0 {
			payload, _, err := c.AdminApi.User(context.Background(), user)
			checkError(err)
			dumpUsersWithZgroup(*payload.User)
			return
		}
		selected, err := selectUser(c, filter)
		checkError(err)
		if selected == nil {
			fmt.Fprintln(os.Stderr, bold("No more results"))
			os.Exit(1)
		}
		dumpUsersWithZgroup(*selected)
	},
}

func init() {
	userCmd.Flags().Int64VarP(&user, "user-id", "u", 0, "user id of user to return")
	userCmd.Flags().StringVarP(&filter, "filter", "f", "", "names of users to filter for")
	userCmd.Flags().BoolVarP(&all, "all", "a", false, "dump all users")

	adminCmd.AddCommand(userCmd)
}
