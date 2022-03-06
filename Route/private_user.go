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
	//用户充值
	group.GET("/user_top_up", user_top_up)
}

//用户信息接口
func UserInfo(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

	new_user, err := Data.Data_Get_userid(strconv.Itoa(user.Id))
	if err != nil {
		r.Response.WriteJson(utils.Get_response_json(1, "获取个人信息错误"))
		return
	}
	r.Session.Set(Config.Session_user, new_user)

	json := gjson.New(nil)
	json.Set("code", 0)
	json.Set("number", new_user.Number)
	json.Set("id", new_user.Id)
	json.Set("money", new_user.Money)
	json.Set("alipay_number", new_user.Alipay_number)
	json.Set("freeze_money", new_user.Freeze_money)

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

//个人发布的单子
func User_fadan(r *ghttp.Request) {
}

//个人接了哪些单子
func User_jiedan(r *ghttp.Request) {
}

//用户充值
func user_top_up(r *ghttp.Request) {
	session_user := r.Session.Get(Config.Session_user)
	user := session_user.(*Bean.User)

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
