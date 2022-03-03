package Route

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"io/ioutil"
	"os"
	"platform/Bean"
	"platform/Config"
	"platform/Data"
	"platform/Route/Filter"
	"platform/log"
	"platform/utils"
	"strings"
)

func init() {
	g := g.Server()
	group := g.Group("/private_user")
	group.Middleware(Filter.Middleware)

	//个人信息
	group.GET("/UserInfo", UserInfo)
	//个人更换信息
	group.GET("/update_userinfo", update_touxiang)
	//个人获取头像
	group.GET("/get_touxiang", Get_touxiang)
	//个人发布的单子
	group.GET("/User_fadan", User_fadan)
	//个人接了哪些单子
	group.GET("/User_jiedan", User_jiedan)
}

//用户信息接口
func UserInfo(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("number", user.Number)
	json.Set("id", user.Id)

	r.Response.WriteJson(json)
}

//用户更新头像接口
func update_touxiang(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	touxiangid := r.GetInt("touxiangid")
	err := Data.Check_fileid(touxiangid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找头像失败"))
		return
	}
	err = Data.Update_User_touxiangid(user, touxiangid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找头像失败"))
		return
	}

}

//用户获取头像接口
func Get_touxiang(r *ghttp.Request) {

	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	filename, err := Data.Get_Img_filename(user.Img)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找文件失败"))
		return
	}

	s := strings.Split(filename, ".")
	if len(s) != 2 {
		r.Response.WriteJson(utils.Get_response_json(1, "查找文件失败"))
		return
	}

	f, err := os.OpenFile(Config.Img_catalog+filename, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找文件失败"))
		log.File_read().Line().Println("文件读取错误", err.Error())
		return
	}

	contentByte, err := ioutil.ReadAll(f)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找文件失败"))
		return
	}

	r.Response.Header().Set("Content-Type", "image/"+utils.File_biaozhun_name[s[1]])
	//r.Response.Header().Set("Accept-Ranges", "bytes")
	r.Response.Header().Set("Content-Disposition", fmt.Sprintf(`attachment;filename="%s"`, "img"))
	r.Response.Write(contentByte)
}

//个人发布的单子
func User_fadan(r *ghttp.Request) {

}

//个人接了哪些单子
func User_jiedan(r *ghttp.Request) {
}
