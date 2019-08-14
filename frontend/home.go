package main

import (
	"fmt"

	"github.com/arschles/vectyshortener/frontend/component"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	uuid "github.com/satori/go.uuid"
	router "marwan.io/vecty-router"
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
					short := uuid.NewV4().String()[0:5]
					h.shortened = short
					vecty.Rerender(h)
				}),
			)),
		),
		elem.Div(vecty.Text(h.shortened)),
		elem.Div(
			router.Link(
				fmt.Sprintf("/created/%s", h.shortened),
				"Create",
				router.LinkOptions{Class: "link"},
			),
		),
	)
}

func shorten() string {
	return uuid.NewV4().String()
}
