package main

import (
	"github.com/arschles/vectyshortener/frontend/style"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
)

func (p *pageView) markdownRow() vecty.ComponentOrHTML {
	return elem.Div(
		style.Class("col-md-4 mb-3"),
		elem.TextArea(
			vecty.Markup(
				vecty.Style("font-family", "monospace"),
				vecty.Property("rows", 14),
				vecty.Property("cols", 70),

				// When input is typed into the textarea, update the local
				// component state and rerender.
				event.Input(func(e *vecty.Event) {
					p.Input = e.Target.Get("value").String()
					vecty.Rerender(p)
				}),
			),
			vecty.Text(p.Input), // initial textarea text.
		),
	)
}
