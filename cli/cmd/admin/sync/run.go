package sync

import (
	"context"
	"fmt"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/print"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run synchronizes user and group state to the stored Google sheet",
	Run: func(cmd *cobra.Command, args []string) {
		client := clientext.NewClient()
		payload, _, err := client.AdminApi.RunSync(context.Background())
		utils.CheckError(err)
		fmt.Println(print.Noticef("Data synchronized to: %s", payload.Sheet))
	},
}
