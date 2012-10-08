package main

import (
    "net/http"
    "io/ioutil"
    _ "fmt"
    _ "encoding/json"
    "mymysql/mysql"
    _ "mymysql/thrsafe"
    "strings"
)

func (this AdminController)IndexAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    this.SetBody("var/html/admin/index.html", nil)
}

// 最初的欢迎页面
func (this AdminIframeController)WelcomeAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    this.SetBody("var/html/admin/welcome.html", p)
}

// 左侧的列表
func (this AdminAjaxController)TreeAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    b, err := ioutil.ReadFile("var/config/tree_" + Config.Get("env") + ".json")
    if (err != nil) {
        panic(err)
    }
    this.OutputRaw(w, r, b)
}

// 密码显示页面
func (this AdminIframeController)PasswordAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    this.SetBody("var/html/admin/password.html", p)
}

func (this AdminAjaxController)PasswordchgAction(w http.ResponseWriter, r *http.Request, p map[string]interface{}) {
    err := r.ParseForm()
    if err != nil {
        this.Output(w, r, 500, "内部错误")
    }
    oldPassword := r.FormValue("oldPassword")
    newPassword := r.FormValue("newPassword")
    checkPassword := r.FormValue("checkPassword")
    if oldPassword == "" || newPassword == "" || checkPassword == "" {
        this.Output(w, r, 100, "密码输入错误")
    }
    
    // 去数据库验证
    if checkPassword != newPassword {
        this.Output(w, r, 100, "密码输入错误")
    }
    db := mysql.New("tcp", "", Config.Get("mysql"), Config.Get("mysql_user"), Config.Get("mysql_password"), Config.Get("mysql_db"))
    
    if err := db.Connect(); err != nil {
        this.Output(w, r, 500, "内部错误")
    }
    defer db.Close()
    
    println("==========this.user============")
    userid := p["UserId"];
    rows, res, err := db.Query("select * from admin_table where admin_id = '%s'", userid)
    if err != nil || len(rows) <= 0 {
        this.Output(w, r, 500, "内部错误")
    }
    
    name := res.Map("admin_password")
    admin_password_db := rows[0].Str(name)
    
    if admin_password_db != oldPassword {
        this.Output(w, r, 100, "密码输入错误")
    }
    update, err := db.Prepare("update admin_table SET admin_password=? where admin_id=?")
    if _,_, err = update.Exec(newPassword, userid); err != nil {
        this.Output(w, r, 500, "内部错误")
    }
    this.Output(w, r, 0, "操作成功")
}