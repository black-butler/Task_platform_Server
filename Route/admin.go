package Route

import (
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

	group.POST("/Get_unreviewed_task", Get_unreviewed_task)
	group.POST("/unreviewed_task", unreviewed_task)
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
