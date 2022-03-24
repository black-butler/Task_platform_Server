package Route

import (
	"platform/Bean"
	"platform/Data"
)

//设置 对方消息未读
func Set_weidu(Work_order *Bean.Work_order, user *Bean.User) {
	if Work_order.Userid == user.Id {
		Data.Data_update_word_unread(Work_order.Id, 1, 0)
	} else if Work_order.Task_userid == user.Id {
		Data.Data_update_word_unread(Work_order.Id, 0, 1)
	}
}
