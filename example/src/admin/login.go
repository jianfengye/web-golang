package main

import (
	"net/http"
    "mymysql/mysql"
    _ "mymysql/thrsafe"
    "utility/web/session"
)

func (this LoginController)IndexAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    isValid, _ := userValid(w, r)
    if isValid {
        this.WebRedirect("/admin/index")
    }
    this.SetBody("var/html/login/index.html", p)
}

func (this LoginController)LoginAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    r.ParseForm()
    admin_name := r.FormValue("admin_name")
    admin_password := r.FormValue("admin_password")
    if admin_name == "" || admin_password == ""{
        this.WebRedirect("/login/index")
    }
    
    // 去数据库中取数据
    db := mysql.New("tcp", "", Config.Get("mysql"), Config.Get("mysql_user"), Config.Get("mysql_password"), Config.Get("mysql_db"))
    
    if err := db.Connect(); err != nil {
        this.WebRedirect("/login/index")
    }
    defer db.Close()
    
    rows, res, err := db.Query("select * from admin_table where admin_name = '%s'", admin_name)
    if err != nil {
        this.WebRedirect("/login/index")
    }

    if len(rows) == 0 {
        this.WebRedirect("/login/index")
    }
    
    name := res.Map("admin_password")
    admin_password_db := rows[0].Str(name)
    if admin_password_db != admin_password {
        this.WebRedirect("/login/index")
    }
    sessions, err := session.SessionStart(r, w)
    if err != nil {
        println(err)
        this.WebRedirect("/login/index")
    }

    sessions.Set("admin_name", rows[0].Str(res.Map("admin_name")))
    sessions.Set("admin_id", rows[0].Str(res.Map("admin_id")))
    println("LoginAction set session")
    this.WebRedirect("/admin/index")
}