package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type created struct {
	vecty.Core
}

func (c *created) Render() vecty.ComponentOrHTML {
	return elem.Div()
}
