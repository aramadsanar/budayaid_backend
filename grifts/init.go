package grifts

import (
	"budayaid/budayaid_backend/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
