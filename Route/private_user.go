package Route

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/guuid"
	"github.com/smartwalle/alipay/v3"
	"io/ioutil"
	"os"
	"platform/Bean"
	"platform/Config"
	"platform/Data"
	"platform/Route/Filter"
	"platform/log"
	"platform/utils"
	"strconv"
	"strings"
	"unicode/utf8"
)

func init() {
	g := g.Server()
	group := g.Group("/private_user")
	group.Middleware(Filter.Middleware)

	//个人信息
	group.POST("/UserInfo", UserInfo)
	//个人更换信息
	group.POST("/update_userinfo", update_touxiang)
	//个人获取头像
	group.GET("/get_touxiang", Get_touxiang)
	//根据用户id获取头像
	group.GET("/get_touxiang_id", Get_touxiang_id)
	//个人发布单子
	group.POST("/User_fadan", User_fadan)
	//个人接了单子
	group.POST("/User_jiedan", User_jiedan)
	//个人接单工单
	group.POST("/User_jiegongdan", User_jiegongdan)
	//个人审核工单
	group.POST("/User_shengongdan", User_shengongdan)
	//工单未读数量
	group.POST("/User_seecount", User_seecount)
	//消息已读
	group.POST("/Message_see_read", Message_see_read)
	//用户充值
	group.POST("/user_top_up", user_top_up)
	//用户申请提现
	group.POST("/apply_for_withdraw_deposit", apply_for_withdraw_deposit)
	//用户提现记录
	group.POST("/withdraw_deposit_record", withdraw_deposit_record)
	//绑定支付宝
	group.POST("/bound_alipay_number", bound_alipay_number)
	//退出登录
	group.POST("/log_out", log_out)
}

//用户信息接口
func UserInfo(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	Data.Data_refre_userid(user)

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("number", user.Number)
	json.Set("id", user.Id)                       //用户id
	json.Set("money", user.Money)                 //用户余额
	json.Set("alipay_number", user.Alipay_number) //用户绑定的支付宝
	json.Set("freeze_money", user.Freeze_money)   //冻结余额
	json.Set("admin", user.Admin)                 //是否是管理员0 不是 1是

	r.Response.WriteJson(json)
}

//用户更新头像接口
func update_touxiang(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	touxiangid := r.GetInt("touxiangid")
	_, err := Data.Data_Get_Img_filename(touxiangid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找头像失败"))
		return
	}
	err = Data.Data_Update_User_touxiangid(user, touxiangid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找头像失败"))
		return
	}

	r.Response.WriteJson(utils.Get_response_json(0, "更换头像成功"))
}

//用户获取头像接口
func Get_touxiang(r *ghttp.Request) {

	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	filename, err := Data.Data_Get_Img_filename(user.Img)
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

//根据用户id获取头像
func Get_touxiang_id(r *ghttp.Request) {

	userid := r.GetString("userid")
	user, err := Data.Data_Get_userid(userid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "查找用户失败"))
		return
	}

	filename, err := Data.Data_Get_Img_filename(user.Img)
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
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	result, err := Data.Data_oneself_publish_tasks(user.Id)
	if err != nil {
		utils.Get_response_json(1, err.Error())
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", result)
	r.Response.WriteJson(json)
}

//个人接了哪些单子
func User_jiedan(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	result, err := Data.Data_oneself_receive_tasks(user.Id)
	if err != nil {
		utils.Get_response_json(1, err.Error())
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", result)
	r.Response.WriteJson(json)
}

//接单工单_消息
func User_jiegongdan(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	result, err := Data.Data_oneself_receive_work_order(user.Id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", result)
	r.Response.WriteJson(json)
}

//审核工单_消息
func User_shengongdan(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	result, err := Data.Data_oneself_publish_work_order(user.Id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, err.Error()))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", result)
	r.Response.WriteJson(json)
}

//获取未读工单
func User_seecount(r *ghttp.Request) {

	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	user_see, taskuser_see, err := Data.Data_gongdan_see_unread(user.Id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取未读数量错误"))
		return
	}

	json := gjson.New(nil)
	json.Set("jiedan_see", user_see)
	json.Set("shenhe_see", taskuser_see)
	r.Response.WriteJson(json)
}

//消息已读
func Message_see_read(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	workid := r.GetInt("wordid")

	err := Data.Data_update_message_read(user.Id, workid)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "消息错误"))
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	r.Response.WriteJson(json)
}

//用户充值
func user_top_up(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	user_suo, err := utils.Get_user_suo(user.Id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "操作用户锁失败"))
		return
	}
	user_suo.Lock()
	defer user_suo.Unlock()

	float_money := r.GetFloat64("money")
	if float_money != float64(int(float_money)) {
		r.Response.WriteJson(utils.Get_response_json(1, "只能充值整数"))
		return
	}

	money := r.GetInt("money")
	if money <= 0 || money > 500 {
		r.Response.WriteJson(utils.Get_response_json(1, "请输入大于0小于500的金额"))
		return
	}

	order_number := guuid.New().String()

	var p = alipay.TradeAppPay{}
	p.NotifyURL = Config.AliPay_NotifyURL
	//p.ReturnURL = "http://xxx"
	p.Subject = "账户充值"
	p.OutTradeNo = order_number
	p.TotalAmount = strconv.Itoa(money)

	url, err := utils.Client.TradeAppPay(p)
	if err != nil {
		log.Alipay_log().Line().Println("创建支付宝订单失败", err.Error())
		r.Response.WriteJson(utils.Get_response_json(1, "创建订单失败"))
		return
	}

	//添加订单
	err = Data.Data_pay_Create_dingdan(user, order_number, money)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "创建订单失败"))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", url)
	r.Response.WriteJson(json)
}

//个人申请提现
func apply_for_withdraw_deposit(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	user_suo, err := utils.Get_user_suo(user.Id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "提现锁失败"))
		return
	}
	user_suo.Lock()
	defer user_suo.Unlock()

	Data.Data_refre_userid(user)

	if user.Alipay_number == "" || user.Alipay_name == "" {
		r.Response.WriteJson(utils.Get_response_json(1, "未绑定支付宝"))
		return
	}

	money := r.GetFloat64("money")
	if money != float64(int(money)) {
		r.Response.WriteJson(utils.Get_response_json(1, "只能提现整数"))
		return
	}

	if int(money) == 0 {
		r.Response.WriteJson(utils.Get_response_json(1, "提现金额错误"))
		return
	}

	if int(money) > user.Money {
		r.Response.WriteJson(utils.Get_response_json(1, "提现金额不足"))
		return
	}

	err = Data.Data_delete_user_money(user.Id, int(money))
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "提现失败"))
		return
	}

	err = Data.Data_Add_withdraw_deposit(user.Id, int(money), user.Alipay_number, user.Alipay_name)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "提现失败"))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", "提现成功")
	r.Response.WriteJson(json)
}

//提现记录
func withdraw_deposit_record(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	result, err := Data.Data_withdraw_deposit_record(user.Id)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "提现记录错误"))
		return
	}

	r.Response.WriteJson(result)
}

//绑定支付宝
func bound_alipay_number(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	alipay_number := r.GetString("alipay_number")
	alipay_name := r.GetString("alipay_name")

	if utf8.RuneCountInString(alipay_number) > 50 || utf8.RuneCountInString(alipay_name) > 10 {
		r.Response.WriteJson(utils.Get_response_json(1, "数据不合法"))
		return
	}

	err := Data.Data_update_user_alipay(user.Id, alipay_number, alipay_name)
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "绑定支付宝失败"))
		return
	}

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("body", "绑定支付宝成功")
	r.Response.WriteJson(json)
}

//登出
func log_out(r *ghttp.Request) {
	r.Session.Remove(Config.Session_user)

	json := gjson.New(nil)
	json.Set("code", "0")
	json.Set("body", "退出成功")
	r.Response.WriteJson(json)
}
