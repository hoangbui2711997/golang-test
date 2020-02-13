package controllers

import "github.com/astaxie/beego"

type BlogsController struct {
	beego.Controller
}

func (this *BlogsController) Get() {
	this.Layout = "layout_blog.tpl"
	this.TplName = "blogs/index.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "blogs/html_head.tpl"
	this.LayoutSections["Scripts"] = "blogs/scripts.tpl"
	this.LayoutSections["Sidebar"] = ""
}