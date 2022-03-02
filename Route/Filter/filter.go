package Filter

import (
	"github.com/gogf/gf/net/ghttp"
	"platform/Bean"
	"platform/Config"
	"platform/utils"
)

//过滤未登录用户
func Middleware(r *ghttp.Request) {
	var user *Bean.User
	err := r.Session.GetStruct(Config.Session_user, user)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取信息失败"))
		return
	}

	if user == nil {
		r.Response.WriteJson(utils.Get_response_json(1, "用户未登录"))
		return
	}

	r.Middleware.Next()
}
