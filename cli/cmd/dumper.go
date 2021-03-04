package cmd

import (
	"os"
	"strconv"
	"time"

	"github.com/GracepointMinistries/hub/client"
	"github.com/olekukonko/tablewriter"
)

func dumpUser(user *client.User, zgroup *client.Zgroup, force bool) {
	table := tablewriter.NewWriter(os.Stdout)
	columns := []string{"ID", "Name", "Email", "Created"}
	headerColors := []tablewriter.Colors{
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
		tablewriter.Colors{tablewriter.Bold},
	}
	columnColors := []tablewriter.Colors{
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor},
		tablewriter.Colors{},
		tablewriter.Colors{},
	}
	values := []string{
		strconv.Itoa(int(user.Id)),
		user.Name,
		user.Email,
		user.CreatedAt.Format(time.RFC1123),
	}
	if zgroup != nil {
		columns = append(columns, "ZGroup")
		values = append(values, zgroup.Name)
		headerColors = append(headerColors, tablewriter.Colors{tablewriter.Bold})
		columnColors = append(columnColors, tablewriter.Colors{})
	} else if force {
		columns = append(columns, "ZGroup")
		values = append(values, "-")
		headerColors = append(headerColors, tablewriter.Colors{tablewriter.Bold})
		columnColors = append(columnColors, tablewriter.Colors{})
	}
	table.SetHeader(columns)
	table.Append(values)
	table.SetHeaderColor(headerColors...)
	table.SetColumnColor(columnColors...)
	table.Render()
}

func dumpZGroupUsers(zgroup *client.Zgroup, users []client.User) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Zgroup",
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
	name := zgroup.Name
	if zgroup.Archived {
		name += " " + warning("(archived)")
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
	table.Render()
}
