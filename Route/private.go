package Route

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/grand"
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
	group := g.Group("/private")
	group.Middleware(Filter.Middleware)

	//首页内容
	group.GET("/announcement", announcement)
	//单子列表
	group.GET("/order_list", order_list)
	//单子详情页
	group.GET("/detail", detail)
	//发单子
	group.GET("/submit", submit)

	//上传图片文件
	group.POST("/UploadFile", UploadFile_Img)
	//获取图片
	group.GET("/Get_Img", Get_Img)
}

//首页内容
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
func UploadFile_Img(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

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
func Get_Img(r *ghttp.Request) {

	imgid := r.GetInt("imgid")
	filename, err := Data.Data_Get_Img_filename(imgid)
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
