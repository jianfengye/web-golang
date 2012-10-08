package controller

import(
	"net/http"
)

type Controller interface {
	PreHandler(w http.ResponseWriter, r *http.Request)
	
	ErrorHandler(w http.ResponseWriter, r *http.Request, err error)

	Run(w http.ResponseWriter, r *http.Request)

	SetHeader(h string, hp map[string]interface{})
	SetAction(a string, ap map[string]interface{})
	SetFooter(f string, fp map[string]interface{})
}