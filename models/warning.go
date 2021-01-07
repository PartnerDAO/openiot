package models

import (
	"time"
)

//Warning 警报
type Warning struct {
	ID            int       //ID
	Wtime         time.Time //报警时间
	Maintaintime  int       //保养时间
	Maintainusers string    //保养人
	Todo          string    //保养事项
}
