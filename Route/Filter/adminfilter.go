package Filter

import (
	"github.com/gogf/gf/net/ghttp"
	"platform/Bean"
	"platform/Config"
	"platform/utils"
)

func Admin_Middleware(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	if session_user == nil {
		r.Response.WriteJson(utils.Get_response_json(1, "用户未登录"))
		return
	}

	user, ok := session_user.(*Bean.User)
	if ok == false {
		r.Response.WriteJson(utils.Get_response_json(1, "用户未登录"))
		return
	}

	if user.Admin == 0 {
		r.Response.WriteJson(utils.Get_response_json(1, "非法访问"))
		return
	}

	r.Middleware.Next()
}
