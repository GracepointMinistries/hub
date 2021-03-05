package sync

import (
	"encoding/json"
)

func (g *groupSlice) String() string {
	data, _ := json.MarshalIndent(g, "", " ")
	return string(data)
}

func (u *userSlice) String() string {
	data, _ := json.MarshalIndent(u, "", " ")
	return string(data)
}
