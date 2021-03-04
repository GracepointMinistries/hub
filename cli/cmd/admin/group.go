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

func selectGroup(c *client.APIClient, filter string) (*client.Zgroup, error) {
	return clientext.PageGroups(c, filter, func(final bool, zgroups []client.Zgroup) (*client.Zgroup, error) {
		lookup, options := print.GroupSelectOptions(final, zgroups)
		prompt := promptui.Select{
			Label:        print.Bold("Select ZGroup"),
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
	Short: "Retrieve information about a zGroup",
	Run: func(cmd *cobra.Command, args []string) {
		c := clientext.NewClient()
		if zgroup == 0 {
			selected, err := selectGroup(c, filter)
			utils.CheckError(err)
			if selected == nil {
				fmt.Fprintln(os.Stderr, print.Bold("No more results"))
				os.Exit(1)
			}
			zgroup = selected.Id
		}
		payload, _, err := c.AdminApi.Zgroup(context.Background(), zgroup)
		utils.CheckError(err)
		print.DumpGroup(payload.Zgroup, payload.Users)
	},
}

func init() {
	groupCmd.Flags().Int64VarP(&zgroup, "group-id", "g", 0, "zGroup to retrieve information for")
	groupCmd.Flags().StringVarP(&filter, "filter", "f", "", "names of zGroups to filter for")
}
