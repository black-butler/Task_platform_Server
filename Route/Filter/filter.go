package Filter

import (
	"github.com/gogf/gf/net/ghttp"
	"platform/Bean"
	"platform/Config"
	"platform/utils"
)

//过滤未登录用户
func Middleware(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	if session_user == nil {
		r.Response.WriteJson(utils.Get_response_json(1, "用户未登录"))
		return
	}

	_, ok := session_user.(*Bean.User)
	if ok == false {
		r.Response.WriteJson(utils.Get_response_json(1, "用户未登录"))
		return
	}

	r.Middleware.Next()
}
