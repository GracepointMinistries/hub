package sync

import (
	"encoding/json"
)

func (i *innerGroup) String() string {
	data, _ := json.MarshalIndent(i, "", " ")
	return string(data)
}

func (g *groupSlice) String() string {
	data, _ := json.MarshalIndent(g, "", " ")
	return string(data)
}

func (i *innerUser) String() string {
	data, _ := json.MarshalIndent(i, "", " ")
	return string(data)
}

func (u *userSlice) String() string {
	data, _ := json.MarshalIndent(u, "", " ")
	return string(data)
}
