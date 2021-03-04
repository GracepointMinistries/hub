package print

import (
	"fmt"

	"github.com/GracepointMinistries/hub/client"
)

var (
	// RetrieveMoreOption is used to page data
	RetrieveMoreOption = Warning("Retrieve more")
	// ExitOption is used to stop paging
	ExitOption = Critical("Exit")
)

// GroupSelectOptions gives a list of options to use in selecting a group
func GroupSelectOptions(final bool, zgroups []client.Zgroup) (map[int]client.Zgroup, []string) {
	lookup := make(map[int]client.Zgroup, len(zgroups))
	options := make([]string, len(zgroups)+1)
	for i, group := range zgroups {
		name := fmt.Sprintf("%s %s", Bold(group.Name), Noticef("[ID: %d]", group.Id))
		if group.Archived {
			name = fmt.Sprintf("%s %s %s", Bold(group.Name), Warning("(archived)"), Noticef("[ID: %d]", group.Id))
		}
		options[i] = name
		lookup[i] = group
	}
	if final {
		options[len(zgroups)] = ExitOption
	} else {
		options[len(zgroups)] = RetrieveMoreOption
	}
	return lookup, options
}

// UserSelectOptions gives a list of options to use in selecting a user
func UserSelectOptions(final bool, users []client.UserWithZgroup) (map[int]client.UserWithZgroup, []string) {
	lookup := make(map[int]client.UserWithZgroup, len(users))
	options := make([]string, len(users)+1)
	for i, user := range users {
		name := fmt.Sprintf("%s %s", Bold(user.Name), Noticef("[ID: %d]", user.Id))
		if user.Blocked {
			name = fmt.Sprintf("%s %s %s", Bold(user.Name), Warning("(blocked)"), Noticef("[ID: %d]", user.Id))
		}
		options[i] = name
		lookup[i] = user
	}
	if final {
		options[len(users)] = ExitOption
	} else {
		options[len(users)] = RetrieveMoreOption
	}
	return lookup, options
}
