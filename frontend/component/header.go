package component

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

func Header(text string) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(prop.ID("standard-header")),
		elem.Heading2(vecty.Text(text)),
	)
}
