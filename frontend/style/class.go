package style

import (
	"strings"

	"github.com/gopherjs/vecty"
)

func Class(values ...string) vecty.MarkupList {
	return vecty.Markup(vecty.Style("class", strings.Join(values, " ")))
}
