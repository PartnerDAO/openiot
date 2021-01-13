package main

import (
	"html/template"
	"net/http"
	"openiot/controllers"
)

func main() {
	//静态资源处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//静态资源处理by页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main", IndexHandler)
	http.ListenAndServe(":9090", nil)
	http.HandleFunc("/login", controllers.Login)
}

//IndexHandler 首页模板函数
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行模板
	t.Execute(w, "")
}
