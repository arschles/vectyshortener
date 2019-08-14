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
					shortened, err := shorten(val)
					if err != nil {
						// what should we do if it fails???
						return
					}
					h.shortened = fmt.Sprintf("rad.shortener.co/%s", shortened)
					vecty.Rerender(h)
				}),
			)),
		),
		elem.Div(vecty.Text(h.shortened)),
	)
}

func shorten(s string) (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return u
}
