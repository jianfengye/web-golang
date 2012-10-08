package controller

import(
	"net/http"
	"reflect"
	"errors"
	"utility/web/util"
	"fmt"
	"encoding/json"
	"strconv"
)

type AjaxController struct {
	Header string
	HeaderParam map[string]interface{}

	Footer string
	FooterParam map[string]interface{}

	Action string
	ActionParam map[string]interface{}

	ParentController interface{}
}

func NewAjaxController() *AjaxController{
	return &AjaxController{}
}

func (this *AjaxController)RegisterController(i interface{}){
	this.ParentController = i
}

func (this *AjaxController)SetHeader(h string, hp map[string]interface{}){
}

func (this *AjaxController)SetAction(a string, ap map[string]interface{}){
	this.Action = a
	this.ActionParam = ap
}

func (this *AjaxController)SetFooter(f string, fp map[string]interface{}){
}

func (this *AjaxController)OutputBody(w http.ResponseWriter, r *http.Request) {
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
    	p = reflect.ValueOf(map[string]string{})
    }
    method.Call([]reflect.Value{responseValue, requestValue, p})
}

func (this *AjaxController)PreHandler(w http.ResponseWriter, r *http.Request) {

}

func (this *AjaxController)ErrorHandler(w http.ResponseWriter, r *http.Request, err error){
	panic(err)
}

func (this *AjaxController)Run(w http.ResponseWriter, r *http.Request) {
	defer func(){
		if err := recover(); err != nil {
			switch t := err.(type) {
			case util.HttpCmd:
				if err.(util.HttpCmd).Command == "redirect" {
					http.Redirect(w, r, err.(util.HttpCmd).Action.(string), http.StatusFound)
					return
				} else if err.(util.HttpCmd).Command == "output" {
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
}

func (this *AjaxController)OutputJson(w http.ResponseWriter, r *http.Request, data interface{}){
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	this.OutputRaw(w, r, b)
}

func (this *AjaxController)OutputRaw(w http.ResponseWriter, r *http.Request, data []byte){
	w.Header().Set("Content-Type", "application/json");
	w.Write(data)
	panic(util.NewHttpCmd("output", nil))
}

func (this *AjaxController)Output(w http.ResponseWriter, r *http.Request, ret int, reason string) {
	this.OutputRaw(w, r, []byte("{\"Ret\": \"" + strconv.Itoa(ret) + "\", \"Reason\" : \"" + reason + "\"}"))
}