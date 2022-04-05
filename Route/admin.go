package Route

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"platform/Bean"
	"platform/Config"
	"platform/Data"
	"platform/Route/Filter"
	"platform/constant"
	"platform/log"
	"platform/utils"
)

func init() {
	g := g.Server()
	group := g.Group("/admin")
	group.Middleware(Filter.Admin_Middleware)

	group.POST("/Get_unreviewed_task", Get_unreviewed_task) //查看未审核任务
	//group.POST("/Get_all_task",Get_all_task)
	group.POST("/unreviewed_task", unreviewed_task)             //审核单子
	group.POST("/update_lun_img", update_lun_img)               //更换轮播图
	group.POST("/withdraw_deposit_list", withdraw_deposit_list) //申请提现列表
	group.POST("/tixain", tixain)                               //处理提现
}

//获取未审核单子
func Get_unreviewed_task(r *ghttp.Request) {
	log.File_admin_log().Println("查看未审核单子")
	result, err := Data.Data_get_all_unreviewed_task()
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取未审核任务失败"))
		return
	}
	r.Response.WriteJson(result)
}

//查看所有状态的任务
//func Get_all_task(r *ghttp.Request){
//	taskid := r.GetInt("taskid")
//	result,err := Data.Data_Get_task_id_all_type(taskid)
//	if err!=nil{
//		r.Response.WriteJson(utils.Get_response_json(1,"获取该任务失败"))
//		return
//	}
//	json := gjson.New(nil)
//	json.Set("code",0)
//	json.Set("body",result)
//}

//审核单子
func unreviewed_task(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)
	taskid := r.GetInt("taskid")
	status := r.GetInt("status")

	log.File_admin_log().Println(user.Id, "管理员审核单子", "taskid:", taskid, "status:", status)

	if !(status == constant.Zhengchang || status == constant.Butongguo) {
		r.Response.WriteJson(utils.Get_response_json(1, "非法操作"))
		return
	}

	err := Data.Data_update_task_status(taskid, status)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "操作任务状态失败"))
		return
	}

	r.Response.WriteJson(utils.Get_response_json(0, "操作成功"))
}

//更换轮播图
func update_lun_img(r *ghttp.Request) {
	zhi := r.GetString("zhi")

	err := Data.Data_update_home_lun_img(zhi)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}
	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", "更换成功")
	r.Response.WriteJson(json)
}

//提现列表
func withdraw_deposit_list(r *ghttp.Request) {
	result, err := Data.Data_withdraw_deposit_apply_for()
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取提现列表失败"))
		return
	}
	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", result)
	r.Response.WriteJson(json)
}

//处理提现
func tixain(r *ghttp.Request) {
	id := r.GetInt("id")

	err := Data.Data_update_deposit_apply_status(id, constant.Tixian_yiwancheng)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "提现失败，数据库错误"))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", "操作成功")
	r.Response.WriteJson(json)
}
