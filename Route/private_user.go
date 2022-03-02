package Route

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"platform/Bean"
	"platform/Config"
	"platform/Route/Filter"
	"platform/utils"
)

func init() {
	g := g.Server()
	group := g.Group("/private_user")
	group.Middleware(Filter.Middleware)

	//个人信息
	group.GET("/UserInfo", UserInfo)
	//个人更换信息
	group.GET("/update_touxiang", update_touxiang)
	//个人发布的单子
	group.GET("/User_fadan", User_fadan)
	//个人接了哪些单子
	group.GET("/User_jiedan", User_jiedan)
}

//用户信息接口
func UserInfo(r *ghttp.Request) {
	var user *Bean.User
	err := r.Session.GetStruct(Config.Session_user, user)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取信息失败"))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("number", user.Number)
	json.Set("id", user.Id)

	r.Response.WriteJson(json)
}

//用户更新头像接口
func update_touxiang(r *ghttp.Request) {

}

//个人发布的单子
func User_fadan(r *ghttp.Request) {
}

//个人接了哪些单子
func User_jiedan(r *ghttp.Request) {
}
