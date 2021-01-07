package models

import (
	"time"
)

//Device 设备字段
type Device struct {
	ID     int       //ID
	Dname  string    //设备名字
	Dtime  time.Time //设备时间
	Dstate int       //设备状态
	Sensor           //传感器
}

//Sensor 传感器字段
type Sensor struct {
	ID        int    //ID
	Sname     string //传感器名称
	Ddataname string //传感器单位，例如温度
	Ddata     int64  //传感器数值
}
