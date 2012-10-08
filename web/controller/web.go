package controller

import(
	"net/http"
	"utility/web/template"
	"reflect"
	"errors"
	"utility/web/util"
	"fmt"
)

type WebController struct {
	Header string
	HeaderParam map[string]interface{}

	Footer string
	FooterParam map[string]interface{}

	Body string
	BodyParam map[string]interface{}

	Action string
	ActionParam map[string]interface{}

	ParentController interface{}
}

func NewWebController() *WebController{
	return &WebController{}
}

func (this *WebController)RegisterController(i interface{}){
	this.ParentController = i
}

func (this *WebController)SetHeader(h string, hp map[string]interface{}){
	this.Header = h
	this.HeaderParam = hp
}

func (this *WebController)SetAction(a string, ap map[string]interface{}){
	this.Action = a
	this.ActionParam = ap
}

func (this *WebController)SetFooter(f string, fp map[string]interface{}){
	this.Footer = f
	this.FooterParam = fp
}

func (this *WebController)SetBody(f string, fp map[string]interface{}){
	this.Body = f
	this.BodyParam = fp
}

func (this *WebController)OutputBody(w http.ResponseWriter, r *http.Request) {
	controller := reflect.ValueOf(this.ParentController)
    method := controller.MethodByName(this.Action)
    if !method.IsValid() {
        panic(errors.New("action does not exist"))
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    var p reflect.Value
    if len(this.ActionParam) > 0 {
    	p = reflect.ValueOf(this.ActionParam)
    } else {
    	p = reflect.ValueOf(map[string]interface{}{})
    }
    method.Call([]reflect.Value{responseValue, requestValue, p})
}

func (this *WebController)PreHandler(w http.ResponseWriter, r *http.Request) {

}

func (this *WebController)ErrorHandler(w http.ResponseWriter, r *http.Request, err error){
	panic(err)
}

func (this *WebController)Run(w http.ResponseWriter, r *http.Request) {
	defer func(){
		if err := recover(); err != nil {
			switch t := err.(type) {
			case util.HttpCmd:
				if err.(util.HttpCmd).Command == "redirect" {
					println("panic redirect:" + err.(util.HttpCmd).Action.(string))
					http.Redirect(w, r, err.(util.HttpCmd).Action.(string), http.StatusFound)
					return
				}
			default:
				fmt.Printf("unexpected type %T", t)
			}
			this.ErrorHandler(w, r, err.(error))
		}
	}()
	this.PreHandler(w,r)

	controller := reflect.ValueOf(this.ParentController)
    method := controller.MethodByName("OutputBody")
    if !method.IsValid() {
        panic(errors.New("action does not exist"))
    }
    requestValue := reflect.ValueOf(r)
    responseValue := reflect.ValueOf(w)
    method.Call([]reflect.Value{responseValue, requestValue})
    
	this.Render(w,r)
}

func (this *WebController)Render(w http.ResponseWriter, r *http.Request) {
	if err := template.RenderHtml(w, this.Header, this.HeaderParam); err != nil {
		panic(err)
	}

	if err := template.RenderHtml(w, this.Body, this.BodyParam); err != nil {
		panic(err)
	}

	if err := template.RenderHtml(w, this.Footer, this.FooterParam); err != nil {
		panic(err)
	}
}

func (this *WebController) WebRedirect(url string) {
	panic(util.NewHttpCmd("redirect", url))
}