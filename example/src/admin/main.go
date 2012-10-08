package main

import (
    "net/http"
    "utility/web/template"
)

func init() {
    template.PreLoad("var")
}

func main() {
    println("main")

    http.Handle("/css/", http.FileServer(http.Dir("var")))
    http.Handle("/js/", http.FileServer(http.Dir("var")))
    http.Handle("/images/", http.FileServer(http.Dir("var")))
    http.Handle("/file/", http.FileServer(http.Dir("var")))

    http.HandleFunc("/favicon.ico",faviconHandler)
    http.HandleFunc("/",routeHandler)
    http.ListenAndServe(Config.Get("listen"), nil)
}