package main

import (
    "net/http"
    "strings"
    "utility/web/controller"
)

func routeHandler(w http.ResponseWriter, r *http.Request) {
    pathInfo := strings.Trim(r.URL.Path, "/")
    parts := strings.Split(pathInfo, "/")
    if parts[0] == "" {
        parts[0] = "login"
    }

    if len(parts) == 1 {
        parts = append(parts, "index")
    }

    if len(parts) == 2 {
        parts = append(parts, "web")
    }

    if len(parts) > 3 {
        panic("invalid url")
    }

    // 有三种可能
    // admin/index
    // admin/index/ajax
    var controller controller.Controller
    switch parts[0] {
    case "admin":
        switch parts[2] {
        case "iframe":
            controller = NewAdminIframeController()
        case "ajax":
            controller = NewAdminAjaxController()
        default:
            controller = NewAdminController()
        }
    case "service":
        switch parts[2] {
        case "ajax":
            controller = NewServiceAjaxController()
        }
        
    case "login":
        controller = NewLoginController()
    }
    action := strings.Title(parts[1]) + "Action"
    controller.SetAction(action, nil)
    controller.Run(w, r)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
    
}