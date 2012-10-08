package main

import(
	. "utility/web/controller"
	"net/http"
	"fmt"
)

type AdminController struct {
	*WebController
}

func NewAdminController() AdminController{
	c := AdminController{ NewWebController() }
	c.SetHeader("var/html/admin/header.html", nil)
	c.SetFooter("var/html/admin/footer.html", nil)
	c.RegisterController(c)
	return c
}

type AdminIframeController struct {
	*WebController
}

func NewAdminIframeController() AdminIframeController {
	c := AdminIframeController{ NewWebController() }
	c.SetHeader("var/html/header_iframe.html", nil)
	c.SetFooter("var/html/footer_iframe.html", nil)
	c.RegisterController(c)
	return c
}

func (this AdminIframeController)OutputBody(w http.ResponseWriter, r *http.Request) {
	if isvalid, user := userValid(w, r); isvalid == true {
		fmt.Println(user)
		println("set BodyParam")
		this.ActionParam = user
	} else {
		this.WebRedirect("/login/index")
	}
	
	this.WebController.OutputBody(w, r)
}

type AdminAjaxController struct {
	*AjaxController
}

func NewAdminAjaxController() AdminAjaxController {
	c := AdminAjaxController{ NewAjaxController() }
	c.RegisterController(c)
	return c
}

func (this AdminAjaxController)OutputBody(w http.ResponseWriter, r *http.Request) {
	if isvalid, user := userValid(w, r); isvalid == true {
		this.ActionParam = make(map[string]interface{})
		this.ActionParam["UserId"] = user["UserId"]
		this.ActionParam["UserName"] = user["UserName"]
	} else {
		this.OutputRaw(w, r, []byte("{'ret':'0'}"))
	}
	this.AjaxController.OutputBody(w, r)
}


type ServiceAjaxController struct {
	*AjaxController
}

func NewServiceAjaxController() ServiceAjaxController {
	c := ServiceAjaxController{ NewAjaxController() }
	c.RegisterController(c)
	return c
}

func (this ServiceAjaxController)OutputBody(w http.ResponseWriter, r *http.Request) {
	if isvalid, user := userValid(w, r); isvalid == true {
		this.ActionParam = make(map[string]interface{})
		this.ActionParam["UserId"] = user["UserId"]
		this.ActionParam["UserName"] = user["UserName"]
	} else {
		this.OutputRaw(w, r, []byte("{'ret':'0'}"))
	}
	this.AjaxController.OutputBody(w, r)
}

type LoginController struct {
	*WebController
}

func NewLoginController() LoginController{
	c := LoginController{ NewWebController() }
	c.SetHeader("var/html/header.html", nil)
	c.SetFooter("var/html/footer.html", nil)
	c.RegisterController(c)
	return c
}
