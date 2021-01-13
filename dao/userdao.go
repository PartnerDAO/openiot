package dao

import (
	"openiot/models"
	"openiot/utils"
)

//CheckUserNameAndPsssword 验证用户名密码
func CheckUserNameAndPsssword(username string, password string) (*models.Users, error) {
	// sql
	sqlStr := "select id ,username,password,email from users where username =? and password=?"
	//执行
	row := utils.DB.QueryRow(sqlStr, username, password)
	//拿到Users 结构体
	user := &models.Users{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//CheckUserName 验证用户名
func CheckUserName(username string) (*models.Users, error) {
	// sql
	sqlStr := "select id ,username,password,email from users where username =? "
	//执行
	row := utils.DB.QueryRow(sqlStr, username)
	//拿到Users 结构体
	user := &models.Users{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

//SaveUser 插入用户信息注册保存
func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username,password,password,email)values(?,?,?)"
	//程序执行
	utils.DB.Exec(sqlStr, username, password, email)
	return nil
}
