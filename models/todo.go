package models

import (
	"time"
)

//Todo 任务字段
type Todo struct {
	ID       int
	Tname    string    //事件内容
	Ttime    time.Time //事件生成时间
	Item     string    // 待办任务事项
	State    int       //事件状态
	Deadline int       //截止日期

}
