package print

import (
	"os"
	"strconv"
	"time"

	"github.com/GracepointMinistries/hub/client"
	"github.com/olekukonko/tablewriter"
)

// DumpGroup dumps the group and the given users as a table
func DumpGroup(group *client.Group, users []client.User) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Group",
		"ID",
		"Name",
		"Email",
		"Created",
	})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	)
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{},
		tablewriter.Colors{},
	)
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	name := group.Name
	if group.Archived {
		name += " " + Warning("(archived)")
	}
	for _, user := range users {
		table.Append([]string{
			name,
			strconv.Itoa(int(user.Id)),
			user.Name,
			user.Email,
			user.CreatedAt.Format(time.RFC1123),
		})
	}
	table.SetRowLine(true)
	table.Render()
}
