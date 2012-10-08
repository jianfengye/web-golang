package main

import (
    "net/http"
    "utility/web/session"
)

func userValid(w http.ResponseWriter, r *http.Request) (bool, map[string]interface{}) {
    session, err := session.SessionStart(r, w)
    if err != nil {
        return false, nil
    }
    userId := session.Get("admin_id")
    userName := session.Get("admin_name")
    if userId == "" || userName == "" {
        return false, nil
    }
    
    return true, map[string]interface{}{"UserId": userId, "UserName" : userName}
}