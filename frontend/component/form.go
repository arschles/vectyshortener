package component

import (
	"github.com/arschles/vectyshortener/frontend/style"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type FormRow []vecty.ComponentOrHTML

func NewFormRow(components ...vecty.ComponentOrHTML) FormRow {
	return components
}

func Form(rows ...FormRow) vecty.ComponentOrHTML {
	rowComponents := make([]vecty.MarkupOrChild, len(rows))
	for i, row := range rowComponents {
		rowComponents[i] = elem.Div(
			style.Class("form-row"),
			row,
		)
	}

	return elem.Form(
		rowComponents...,
	)
}
