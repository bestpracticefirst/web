package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"strconv"
	"time"
	"crypto/md5"
	"io"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//用MD5获取唯一值，把这个值放到服务器里
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		age,err:=strconv.Atoi(r.Form.Get("age"))
		if err!=nil{
			fmt.Fprintf(w, "The age is not a number")
			return
		}
		//读取token
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form.Get("username"))
		fmt.Println("password:", r.Form["password"])
		fmt.Println("age:",age)
		if isSelected(r){
			fmt.Println("fruit",r.Form.Get("fruit"))
		}
		if checkedSex(r){
			fmt.Println("sex:",r.Form.Get("gender"))
		}
		template.HTMLEscape(w, []byte(token)) //输出到客户端
	}
}
func checkedSex( r *http.Request) bool{
	slice:=[]int{1,2}
	for _, v := range slice {
		gender,_:=strconv.Atoi(r.Form.Get("gender"));
		if v ==gender {
			return true
		}
	}
	return false
}
func isSelected(r *http.Request) bool{
	slice:=[]string{"apple","pear","banane"}
	r.ParseForm()
	v := r.Form.Get("fruit")
	for _, item := range slice {
		if item == v {
			return true
		}
	}

	return false
}
func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}