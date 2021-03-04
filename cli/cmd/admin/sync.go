package admin

import (
	"context"
	"fmt"
	"os"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/print"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Initialize a new Google sheet for synchronization",
	Run: func(cmd *cobra.Command, args []string) {
		client := clientext.NewClient()
		settings, _, err := client.AdminApi.CurrentSettings(context.Background())
		utils.CheckError(err)
		if settings.Sheet != "" {
			fmt.Println(print.Criticalf("Continuing this operation will stop syncing from:\n\t%s\n", print.Notice(settings.Sheet)))
			prompt := promptui.Prompt{
				Label:     "Are you sure you want to continue",
				IsConfirm: true,
			}
			_, err = prompt.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, print.Warning("Initialization canceled"))
				return
			}
		}
		payload, _, err := client.AdminApi.InitializeSync(context.Background())
		utils.CheckError(err)
		fmt.Println(print.Noticef("Google Sheet created at: %s", payload.Sheet))
	},
}
