package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router /user/list/ [get]
func (c *MainController) List() {
	c.Ctx.WriteString("list")
}

// @router /user/:id [get]
func (c *MainController) GetOne() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("get : " + id)
}

// @router /user/ [post]
func (c *MainController) Insert() {
	c.Ctx.WriteString("insert")
}

// @router /user/:id [put]
func (c *MainController) Update() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("update : " + id)
}

// @router /user/:id [delete]
func (c *MainController) Delete() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("delete : " + id)
}
