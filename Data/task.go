package Data

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/log"
	"platform/utils"
	"strconv"
	"time"
)

//添加任务
func Data_add_task(user *Bean.User, body string, img string, one_money int, sum int) error {
	_, err := g.DB().Model("tasks").Data(g.Map{"userid": user.Id, "body": body, "imgs": img, "one_money": one_money, "sum": sum, "time": time.Now().Format(utils.Time_Format)}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("添加任务", err.Error())
		return errors.New("添加任务失败")
	}
	return nil
}

//根据id获取某个任务
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
func Data_Set_work_order(user *Bean.User, taskid int, finish_time string) error {
	_, err := g.DB().Model("work_order").Data(g.Map{"userid": user.Id, "taskid": taskid, "create_time": time.Now().Format(utils.Time_Format), "finish_time": finish_time}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("添加任务接单记录失败", err.Error())
		return errors.New("添加任务接单记录失败")
	}
	return nil
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
