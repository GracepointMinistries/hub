package sync

import (
	"fmt"
	"strconv"
)

func stringNil(v *string) string {
	if v == nil {
		return "nil"
	}
	return "'" + *v + "'"
}

func intNil(v *int) string {
	if v == nil {
		return "nil"
	}
	return strconv.Itoa(*v)
}

func boolNil(v *bool) string {
	if v == nil {
		return "nil"
	}
	if *v {
		return "true"
	}
	return "false"
}

func (g *groupSlice) String() string {
	s := "["
	for _, group := range *g {
		s += fmt.Sprintf(`
  group {
    ID: %s
    Name: %s
    ZoomLink: %s
    Published: %s
    Archived: %s
  },`, intNil(group.ID), stringNil(group.Name), stringNil(group.ZoomLink), boolNil(group.Published), boolNil(group.Archived))
	}
	s += "\n]"
	return s
}

func (u *userSlice) String() string {
	s := "["
	for _, user := range *u {
		s += fmt.Sprintf(`
  user {
    ID: %s
    Name: %s
    Email: %s
    Blocked: %s
    Group: %s
  },`, intNil(user.ID), stringNil(user.Name), stringNil(user.Email), boolNil(user.Blocked), stringNil(user.Group))
	}
	s += "\n]"
	return s
}
