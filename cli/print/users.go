package print

import (
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/GracepointMinistries/hub/client"
	"github.com/olekukonko/tablewriter"
)

// DumpUsers prints users sorted by group
func DumpUsers(users ...client.UserWithZgroup) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ZGroup", "ID", "Name", "Email", "Created"})
	data := make([][]string, len(users))
	for i, user := range users {
		values := []string{}
		if user.Zgroup != nil {
			name := user.Zgroup.Name
			if user.Zgroup.Archived {
				name += " " + Warning("(archived)")
			}
			values = append(values, name)
		} else {
			values = append(values, "-")
		}
		values = append(
			values,
			strconv.Itoa(int(user.Id)),
			user.Name,
			user.Email,
			user.CreatedAt.Format(time.RFC1123),
		)
		data[i] = values
	}
	// sort by zgroup to keep all data together
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})

	table.AppendBulk(data)
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
	table.SetRowLine(true)
	table.Render()
}
