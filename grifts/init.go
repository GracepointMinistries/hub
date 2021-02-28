package grifts

import (
	"github.com/GracepointMinistries/hub/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
