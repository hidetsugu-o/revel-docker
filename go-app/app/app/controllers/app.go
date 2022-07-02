package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.ViewArgs["test"] = "Hello World"
	c.ViewArgs["test2"] = "こんにちは"
	return c.Render()
}
