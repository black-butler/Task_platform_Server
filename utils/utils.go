package utils

import (
	"platform/Bean"
	"platform/Data"
)

const (
	Time_Format = "2006-01-02 15:04:05"
)

//判断某个float64是不是整数
func Check_float64_zheng(zhi float64) bool {
	if zhi == float64(int(zhi)) {
		return true
	} else {
		return false
	}
}

//设置 对方消息未读
func Set_weidu(Work_order *Bean.Work_order, user *Bean.User) {
	if Work_order.Userid == user.Id {
		Data.Data_update_word_unread(Work_order.Id, 1, 0)
	} else if Work_order.Task_userid == user.Id {
		Data.Data_update_word_unread(Work_order.Id, 0, 1)
	}
}
