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
	"time"
	"unicode/utf8"
)

func init() {
	g := g.Server()
	group := g.Group("/private")
	group.Middleware(Filter.Middleware)

	//首页内容
	group.POST("/announcement", announcement)
	//任务列表
	group.POST("/order_list", order_list)
	//任务详情页
	group.POST("/detail", detail)
	//提交任务
	group.POST("/submit", submit)
	//接任务
	group.POST("/receive_task", receive_task)
	//下架任务

	//上传图片文件
	group.POST("/UploadFile", UploadFile_Img)
	//获取图片
	group.POST("/Get_Img", Get_Img)
}

//首页内容
func announcement(r *ghttp.Request) {
	lunbo_imgs, gonggao, err := Data.Data_Get_home()
	if err != nil {
		return
	}

	json := gjson.New(nil)
	json.Set("lunbo_imgs", lunbo_imgs)
	json.Set("gonggao", gonggao)
	r.Response.WriteJson(json)
}

//单子列表
func order_list(r *ghttp.Request) {
	result, err := Data.Data_get_all_task()
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	json := gjson.New(nil)
	json.Set("code", "0")
	json.Set("body", result)
	r.Response.WriteJson(json)
}

//单子详情页接口
func detail(r *ghttp.Request) {
	id := r.GetInt("id")
	record, err := Data.Data_get_task(id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	json := gjson.New(nil)
	json.Set("code", "0")
	json.Set("body", record)
	r.Response.WriteJson(json)
}

//提交单子接口
func submit(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	title := r.GetString("title")        //任务标题
	TaskCount := r.GetInt("TaskCount")   //任务数量
	price := r.GetInt("price")           //任务单价
	Time_limit := r.GetInt("Time_limit") //任务限时
	endDate := r.GetInt64("endDate")     //超时时间戳
	body := r.GetString("body")          //任务完成流程
	proof := r.GetString("proof")        //任务需要提交的资料
	img := r.GetString("imgs")           //多个图片

	if Time_limit < 10 || Time_limit > 1000 {
		r.Response.WriteJson(utils.Get_response_json(1, "任务限时不能小于10分钟或大于1000分钟"))
		return
	}
	if utf8.RuneCountInString(title) < 5 || utf8.RuneCountInString(title) > 15 {
		r.Response.WriteJson(utils.Get_response_json(1, "任务标题不能大于15小于5个字符"))
		return
	}
	if TaskCount < 1 || TaskCount > 1000 {
		r.Response.WriteJson(utils.Get_response_json(1, "任务数量不能小于1大于1000"))
		return
	}
	if price < 1 || price > 1000 {
		r.Response.WriteJson(utils.Get_response_json(1, "任务单价不能小于1大于1000"))
		return
	}
	now := time.Unix(endDate, 0)
	if time.Now().Unix() >= now.Unix() {
		r.Response.WriteJson(utils.Get_response_json(1, "时间错误"))
		return
	}
	if utf8.RuneCountInString(body) <= 10 || utf8.RuneCountInString(body) > 1000 {
		r.Response.WriteJson(utils.Get_response_json(1, "任务流程不能小于10大于1000"))
		return
	}
	if utf8.RuneCountInString(proof) <= 10 || utf8.RuneCountInString(proof) > 1000 {
		r.Response.WriteJson(utils.Get_response_json(1, "任务需要提交的资料不能小于10大于1000"))
		return
	}

	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	//刷新
	Data.Data_refre_userid(user)

	//任务总金额
	Zong_money := price * TaskCount
	//判断金额是否足够
	if Zong_money > user.Money {
		r.Response.WriteJson(utils.Get_response_json(1, "用户余额不足，请先充值"))
		return
	}

	//校验图片
	imgs := strings.Split(img, ",")
	if len(imgs) > 100 {
		r.Response.WriteJson(utils.Get_response_json(1, "图片过多"))
		return
	}
k:
	for i, v := range imgs {
		if v == "" {
			imgs = append(imgs[:i], imgs[i+1:]...)
			goto k
		}
	}

	status, err := Data.Data_Check_img_ids(imgs)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "图片提交失败"))
		return
	}
	if status == false {
		r.Response.WriteJson(utils.Get_response_json(1, "图片提交失败"))
		return
	}

	//提交任务
	err = Data.Data_add_task(user, title, body, proof, img, TaskCount, price, Time_limit, now.Format(utils.Time_Format))
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "提交任务失败"))
		return
	}

	//更新用户余额
	err = Data.Data_transfer_money(user, Zong_money)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 1)
	json.Set("body", "提交成功")
	r.Response.WriteJson(json)
}

//用户接任务
func receive_task(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	taskid := r.GetInt("taskid")
	task, err := Data.Data_Get_task_id(taskid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取任务失败"))
		return
	}

	word_task, err := Data.Data_Check_user_receive_task(user, task.Id)
	if word_task != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "已接过此任务"))
		return
	}

	finish_time := time.Now().Add(time.Minute * time.Duration(task.Time_limit))
	err = Data.Data_Set_work_order(user, task.Id, task.Userid, finish_time.Format(utils.Time_Format))
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "接任务失败"))
		return
	}

	r.Response.WriteJson(utils.Get_response_json(0, "接任务成功，请在指定时间内完成"))
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
	//json.Set("body","上传成功")

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
