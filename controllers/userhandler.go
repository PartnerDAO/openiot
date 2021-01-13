package controllers

import (
	"html/template"
	"net/http"
	"openiot/dao"
)

//Login 处理用户登陆函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取表单中用户名密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	//调用dao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPsssword(username, password)

	if user.ID > 0 {
		//如果用户名密码正确
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//如果用户名密码错误
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}

}
