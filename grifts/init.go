package grifts

import (
	"github.com/StackFocus/SwagIP-go/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
