package Data

import (
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/constant"
	"platform/log"
	"platform/utils"
	"strconv"
	"time"
)

//添加任务
func Data_add_task(user *Bean.User, title string, body string, audit string, img string, one_money int, freeze_money int, sum int, time_limit int, endDate string) (int64, error) {
	id, err := g.DB().Model("tasks").Data(g.Map{"userid": user.Id, "title": title, "body": body, "audit": audit, "imgs": img, "one_money": one_money, "freeze_money": freeze_money, "sum": sum, "time_limit": time_limit, "endDate": endDate, "time": time.Now().Format(utils.Time_Format)}).InsertAndGetId()
	if err != nil {
		log.Sql_log().Line().Println("添加任务", err.Error())
		return 0, errors.New("添加任务失败")
	}
	return id, nil
}

//根据id获取某个任务 所有任务
func Data_Get_task_id(taskid int) (*Bean.Task, error) {
	task := new(Bean.Task)
	err := g.DB().Model("tasks").Where("id", taskid).Struct(task)
	if err != nil {
		log.Sql_log().Line().Println("根据id获取某个任务", err.Error())
		return nil, errors.New("根据id获取某个任务")
	}

	user, err := Data_Get_userid(strconv.Itoa(task.Userid))
	if err != nil {
		return nil, err
	}

	task.User = user

	return task, nil
}

//添加任务接单记录
func Data_Set_work_order(user *Bean.User, taskid int, taskUserid int, finish_time string) (int64, error) {
	id, err := g.DB().Model("work_order").Data(g.Map{"userid": user.Id, "taskid": taskid, "task_userid": taskUserid, "create_time": time.Now().Format(utils.Time_Format), "finish_time": finish_time}).InsertAndGetId()
	if err != nil {
		log.Sql_log().Line().Println("添加任务接单记录失败", err.Error())
		return 0, errors.New("添加任务接单记录失败")
	}
	return id, nil
}

//获取某个用户对某个任务的接单记录
func Data_Check_user_receive_task(user *Bean.User, taskid int) (*Bean.Work_order, error) {
	Work_order := new(Bean.Work_order)
	err := g.DB().Model("work_order").Where("userid", user.Id).Where("taskid", taskid).Struct(Work_order)
	if err != nil {
		log.Sql_log().Line().Println("获取某个用户对某个任务的接单记录", err.Error())
		return nil, errors.New("获取任务记录失败")
	}

	Work_order.User = user

	task, err := Data_Get_task_id(Work_order.Taskid)
	if err != nil {
		return nil, err
	}
	Work_order.Task = task

	return Work_order, nil
}

//获取某个任务下的所有接单记录
func Get_Task_work_order_all(task *Bean.Task) ([]*Bean.Work_order, error) {
	Work_orders := make([]*Bean.Work_order, 0)
	err := g.DB().Model("work_order").Where("taskid", task.Id).Structs(&Work_orders)
	if err != nil {
		log.Sql_log().Line().Println("获取某个任务下的所有接单记录", err.Error())
		return nil, errors.New("获取任务失败")
	}

	for _, v := range Work_orders {
		v.Task = task
		user, err := Data_Get_userid(strconv.Itoa(v.Userid))
		if err != nil {
			return nil, errors.New("查找用户失败")
		}
		v.User = user
	}
	return Work_orders, nil
}

//获取当前所有任务
func Data_get_all_task() (gdb.Result, error) {
	result, err := g.DB().Model("tasks").Where("status", constant.Zhengchang).All()
	if err != nil {
		log.Sql_log().Line().Println("获取当前所有任务", err.Error())
		return nil, errors.New("获取推荐任务失败")
	}
	return result, nil
}

//根据id获取单个任务 正常状态的任务
func Data_get_task(id int) (gdb.Record, error) {
	Record, err := g.DB().Model("tasks").Where("status", constant.Zhengchang).Where("id", id).One()
	if err != nil {
		log.Sql_log().Line().Println("获取当前所有任务", err.Error())
		return nil, errors.New("查看当前任务失败")
	}
	return Record, nil
}

//获取某个任务的所有接单数量
func Data_get_task_dan_count(taskid int) (int, error) {
	reslut, err := g.DB().Model("work_order").Where("taskid", taskid).WhereNotIn("status", g.Slice{constant.Chaoshi}).All()
	if err != nil {
		log.Sql_log().Line().Println("获取某个任务的所有接单数量", err.Error())
		return 0, errors.New("查看当前任务失败")
	}

	return len(reslut), nil
}

//提交任务资料
func Data_Add_message(user *Bean.User, work_orderid int64, task *Bean.Task, body string, imgs string) error {
	_, err := g.DB().Model("message").Data(g.Map{"userid": user.Id, "workid": work_orderid, "taskid": task.Id, "taskuserid": task.Userid, "body": body, "imgs": imgs, "time": time.Now().Format(utils.Time_Format)}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("提交任务资料", err.Error())
		return errors.New("提交任务资料")
	}
	return nil
}

//获取某个工单对应的全部消息
func Data_get_all_message(Work_order *Bean.Work_order) (gdb.Result, error) {
	result, err := g.DB().Model("message").Where("workid", Work_order.Id).All()
	if err != nil {
		log.Sql_log().Line().Println("获取某个工单对应的全部消息", err.Error())
		return nil, errors.New("获取消息失败")
	}
	return result, nil
}

//更新任务状态
func Data_update_task_status(taskid int, status int) error {
	_, err := g.DB().Model("tasks").Data(g.Map{"status": status}).Where("id", taskid).Update()
	if err != nil {
		log.Sql_log().Line().Println("添加用户余额失败", err.Error())
		return errors.New("更新任务状态失败")
	}
	return nil
}

//更新work接单状态
func Data_update_work_status(work *Bean.Work_order, status int) error {
	_, err := g.DB().Model("work_order").Data(g.Map{"status": status}).Where("id", work.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("添加用户余额失败", err.Error())
		return errors.New("更新接单状态失败")
	}
	return nil
}

//根据id获取工单
func Data_get_Work_orderid(id int) (*Bean.Work_order, error) {

	Work_order := new(Bean.Work_order)
	err := g.DB().Model("work_order").Where("id", id).Struct(Work_order)
	if err != nil {
		log.Sql_log().Line().Println("获取某个用户对某个任务的接单记录", err.Error())
		return nil, errors.New("获取任务记录失败")
	}

	user, err := Data_Get_userid(strconv.Itoa(Work_order.Userid))
	if err != nil {
		return nil, err
	}
	Work_order.User = user

	task, err := Data_Get_task_id(Work_order.Taskid)
	if err != nil {
		return nil, err
	}
	Work_order.Task = task

	return Work_order, nil
}

//扣除任务剩余冻结余额
func Data_delete_task_freeze_money(task *Bean.Task, money int) error {
	_, err := g.DB().Model("tasks").Data(g.Map{"freeze_money": gdb.Raw("freeze_money-" + strconv.Itoa(money))}).Where("id", task.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("扣除任务剩余冻结余额", err.Error())
		return errors.New("任务余额操作失败")
	}
	return nil
}

//查看自己接的任务
func Data_oneself_receive_tasks(userid int) (gdb.Result, error) {
	result, err := g.DB().Model("work_order").Where("userid", userid).All()
	if err != nil {
		log.Sql_log().Line().Println("查看自己接的任务-1:", err.Error())
		return nil, errors.New("查看自己接的任务失败")
	}

	taskids := make([]int, 0)
	for _, v := range result {
		taskids = append(taskids, v["taskid"].Int())
	}

	word_result, err := g.DB().Model("tasks").Where("id", taskids).All()
	if err != nil {
		log.Sql_log().Line().Println("查看自己接的任务-2:", err.Error())
		return nil, errors.New("查看自己接的任务失败")
	}

	return word_result, nil
}

//查看自己发布的任务
func Data_oneself_publish_tasks(userid int) (gdb.Result, error) {
	word_result, err := g.DB().Model("tasks").Where("userid", userid).All()
	if err != nil {
		log.Sql_log().Line().Println("查看自己发布的任务:", err.Error())
		return nil, errors.New("查看自己发布的任务失败")
	}

	return word_result, nil
}

//查看接单工单
//func
