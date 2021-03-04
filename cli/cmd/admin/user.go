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

func selectUser(c *client.APIClient, filter string) (*client.UserWithGroup, error) {
	return clientext.PageUsers(c, filter, func(final bool, users []client.UserWithGroup) (*client.UserWithGroup, error) {
		lookup, options := print.UserSelectOptions(final, users)
		prompt := promptui.Select{
			Label:        print.Bold("Select User"),
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
		if user, found := lookup[index]; found {
			return &user, nil
		}
		return nil, nil
	})
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Retrieve information about a user",
	Run: func(cmd *cobra.Command, args []string) {
		c := clientext.NewClient()
		if all {
			users, err := clientext.AllUsers(c, filter)
			utils.CheckError(err)
			print.DumpUsers(users...)
			return
		}
		if user != 0 {
			payload, _, err := c.AdminApi.User(context.Background(), user)
			utils.CheckError(err)
			print.DumpUsers(*payload.User)
			return
		}
		selected, err := selectUser(c, filter)
		utils.CheckError(err)
		if selected == nil {
			fmt.Fprintln(os.Stderr, print.Bold("No more results"))
			os.Exit(1)
		}
		print.DumpUsers(*selected)
	},
}

func init() {
	userCmd.Flags().Int64VarP(&user, "user-id", "u", 0, "user id of user to return")
	userCmd.Flags().StringVarP(&filter, "filter", "f", "", "names of users to filter for")
	userCmd.Flags().BoolVarP(&all, "all", "a", false, "dump all users")
}
