package Route

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/grand"
	"platform/Bean"
	"platform/Config"
	"platform/Data"
	"platform/log"
	"platform/utils"
	"strings"
)

func init() {
	g := g.Server()
	group := g.Group("/private")

	//公告接口
	group.GET("/announcement", announcement)
	//单子列表
	group.GET("/order_list", order_list)
	//单子详情页
	group.GET("/detail", detail)
	//发单子
	group.GET("/submit", submit)
	//上传图片
	group.GET("/uploading_img", uploading)
	//用户充值
	group.GET("/user_top_up", user_top_up)

	//上传文件
	group.POST("/UploadFile", UploadFile)
}

//公告接口
func announcement(r *ghttp.Request) {

}

//单子列表
func order_list(r *ghttp.Request) {

}

//单子详情页接口
func detail(r *ghttp.Request) {

}

//提交单子接口
func submit(r *ghttp.Request) {

}

//上传图片
func uploading(r *ghttp.Request) {

}

//用户充值
func user_top_up(r *ghttp.Request) {}

//上传文件
func UploadFile(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {
		r.Response.WriteJson(utils.Get_response_json(1, "文件不存在"))
		return
	}
	//检查文件后缀
	houzhui := strings.Split(file.FileHeader.Filename, ".")
	if len(houzhui) != 2 {
		r.Response.WriteJson(utils.Get_response_json(1, "文件错误"))
		return
	}
	file.FileHeader.Filename = grand.S(15) + "." + houzhui[1]
	//检查文件格式
	f, err := file.Open()
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "文件错误"))
		return
	}
	bytes := make([]byte, 10)
	f.Read(bytes)

	var file_type = utils.GetFileType(bytes)
	if !utils.Check_if_img(file_type) {
		r.Response.WriteJson(utils.Get_response_json(1, "文件错误"))
		return
	}

	//保存文件
	filename, err := file.Save(Config.Img_catalog)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "上传文件错误"))
		log.UploadFile_log().Println("文件上传失败", err.Error())
		return
	}

	user := new(Bean.User)
	user.Id = 1
	id, err := Data.Data_Save_file(user, filename)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "上传文件错误"))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("id", id)
	r.Response.WriteJson(json)
}

//获取图片
