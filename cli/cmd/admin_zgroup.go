package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/GracepointMinistries/hub/client"
	"github.com/antihax/optional"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	zgroup int64
	filter string
)

// printers
var (
	warning   = color.New(color.Bold, color.FgYellow).SprintFunc()
	warningf  = color.New(color.Bold, color.FgYellow).SprintfFunc()
	critical  = color.New(color.Bold, color.FgRed).SprintFunc()
	criticalf = color.New(color.Bold, color.FgRed).SprintfFunc()
	notice    = color.New(color.Bold, color.FgGreen).SprintFunc()
	noticef   = color.New(color.Bold, color.FgGreen).SprintfFunc()
	bold      = color.New(color.Bold).SprintFunc()
	boldf     = color.New(color.Bold).SprintfFunc()

	retrieveMoreOption = warning("Retrieve more")
	exitOption         = critical("Exit")
)

func pageZGroups(c *client.APIClient, filter string, handler func(bool, []client.Zgroup) (*client.Zgroup, error)) (*client.Zgroup, error) {
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

func zgroupOptions(final bool, zgroups []client.Zgroup) (map[int]client.Zgroup, []string) {
	lookup := make(map[int]client.Zgroup, len(zgroups))
	options := make([]string, len(zgroups)+1)
	for i, group := range zgroups {
		name := fmt.Sprintf("%s %s", bold(group.Name), noticef("[ID: %d]", group.Id))
		if group.Archived {
			name = fmt.Sprintf("%s %s %s", bold(group.Name), warning("(archived)"), noticef("[ID: %d]", group.Id))
		}
		options[i] = name
		lookup[i] = group
	}
	if final {
		options[len(zgroups)] = exitOption
	} else {
		options[len(zgroups)] = retrieveMoreOption
	}
	return lookup, options
}

func selectZGroup(c *client.APIClient, filter string) (*client.Zgroup, error) {
	return pageZGroups(c, filter, func(final bool, zgroups []client.Zgroup) (*client.Zgroup, error) {
		lookup, options := zgroupOptions(final, zgroups)
		prompt := promptui.Select{
			Label:        bold("Select ZGroup"),
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
		if group, found := lookup[index]; found {
			return &group, nil
		}
		return nil, nil
	})
}

var zgroupCmd = &cobra.Command{
	Use:   "zgroup",
	Short: "Retrieve information about a zGroup",
	Run: func(cmd *cobra.Command, args []string) {
		c := newClient()
		if zgroup == 0 {
			selected, err := selectZGroup(c, filter)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			if selected == nil {
				fmt.Fprintln(os.Stderr, "No more results")
				os.Exit(1)
			}
			zgroup = selected.Id
		}
		payload, _, err := c.AdminApi.Zgroup(context.Background(), zgroup)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
		dumpZGroupUsers(payload.Zgroup, payload.Users)
	},
}

func init() {
	zgroupCmd.Flags().Int64VarP(&zgroup, "zgroup-id", "z", 0, "zGroup to retrieve information for")
	zgroupCmd.Flags().StringVarP(&filter, "filter", "f", "", "names of zGroups to filter for")

	adminCmd.AddCommand(zgroupCmd)
}
