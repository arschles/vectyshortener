package main

import (
	"fmt"

	"github.com/arschles/vectyshortener/frontend/component"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	uuid "github.com/satori/go.uuid"
)

type home struct {
	vecty.Core
	shortened string
}

func (h *home) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Style("text-align", "center")),
		component.Header("Let's Create Some Short Links!"),
		elem.Form(
			elem.Input(vecty.Markup(
				event.Input(func(e *vecty.Event) {
					val := e.Target.Get("value").String()
					h.shortened = fmt.Sprintf("rad.shortener.co/%s", shorten(val)[0:5])
					vecty.Rerender(h)
				}),
			)),
		),
		elem.Div(vecty.Text(h.shortened)),
	)
}

func shorten(s string) string {
	return uuid.NewV4().String()
}
