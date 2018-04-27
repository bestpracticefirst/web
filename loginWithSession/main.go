package main

import (
	"net/http"
	"html/template"
	"web/session"
	"log"
	"time"
)
var globalSessions=session.NewSessionManager()
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.BeginSession(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("loginWithSession/login.html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}
func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.BeginSession(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.Destroy(w,r)
		sess = globalSessions.BeginSession(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("loginWithSession/count.html")
	t.Execute(w, sess.Get("countnum"))
}
func main() {
	http.HandleFunc("/login", login)         //设置访问的路由
	http.HandleFunc("/count", count)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
