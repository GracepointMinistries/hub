package admin

import (
	"context"
	"fmt"

	"github.com/GracepointMinistries/hub/cli/clientext"
	"github.com/GracepointMinistries/hub/cli/print"
	"github.com/GracepointMinistries/hub/cli/utils"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export user and group state to the stored Google sheet",
	Run: func(cmd *cobra.Command, args []string) {
		client := clientext.NewClient()
		payload, _, err := client.AdminApi.DumpSync(context.Background())
		utils.CheckError(err)
		fmt.Println(print.Noticef("Data dumped to: %s", payload.Sheet))
	},
}
