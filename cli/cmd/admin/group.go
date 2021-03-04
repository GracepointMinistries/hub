package admin

import (
	"context"
	"fmt"
	"os"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/print"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/GracepointMinistries/hub/client"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func selectGroup(c *client.APIClient, filter string) (*client.Group, error) {
	return clientext.PageGroups(c, filter, func(final bool, zgroups []client.Group) (*client.Group, error) {
		lookup, options := print.GroupSelectOptions(final, zgroups)
		prompt := promptui.Select{
			Label:        print.Bold("Select Group"),
			Items:        options,
			HideSelected: true,
			Stdout:       utils.NewBellSkipper(),
		}
		index, result, err := prompt.Run()
		if err != nil {
			return nil, err
		}
		if result == print.ExitOption || result == print.RetrieveMoreOption {
			return nil, nil
		}
		if group, found := lookup[index]; found {
			return &group, nil
		}
		return nil, nil
	})
}

var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Retrieve information about a group",
	Run: func(cmd *cobra.Command, args []string) {
		c := clientext.NewClient()
		if group == 0 {
			selected, err := selectGroup(c, filter)
			utils.CheckError(err)
			if selected == nil {
				fmt.Fprintln(os.Stderr, print.Bold("No more results"))
				os.Exit(1)
			}
			group = selected.Id
		}
		payload, _, err := c.AdminApi.Group(context.Background(), group)
		utils.CheckError(err)
		print.DumpGroup(payload.Group, payload.Users)
	},
}

func init() {
	groupCmd.Flags().Int64VarP(&group, "group-id", "g", 0, "group to retrieve information for")
	groupCmd.Flags().StringVarP(&filter, "filter", "f", "", "names of groups to filter for")
}
