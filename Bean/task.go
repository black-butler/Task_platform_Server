package Bean

import "time"

type Work_order struct {
	Id          int
	Userid      int    //接任务用户id
	User        *User  //接任务用户
	Taskid      int    //任务id
	task_userid int    //发布任务用户id
	Task        *Task  //任务对象
	Create_time string //创建时间
	Finish_time string //结束时间
	Status      int
}

//获取当前的时间戳
func (this *Work_order) Create_time_f() time.Time {
	return time.Now()
}

type Task struct {
	Id           int    //任务id
	Userid       int    //发布者id
	User         *User  //发布者对象
	Title        string //任务标题
	Body         string //任务内容
	Audit        string //审核提交内容
	Imgs         string //任务图片
	One_money    int    //单个金额
	freeze_money int    //共冻结的金额
	Sum          int    //任务派单数量
	EndDate      string //任务结束时间
	Time_limit   int    //任务限时多少分钟
	Status       int    //0正常 1下架
	Time         string //创建时间
}
