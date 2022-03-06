package Route

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"platform/Data"
	"platform/Tripartite_api"
	"platform/log"
	"platform/utils"
)

func init() {
	g := g.Server()
	group := g.Group("/Alipay")

	group.POST("/NotifyURL", Alipay_NotifyURL)
}

func Alipay_NotifyURL(r *ghttp.Request) {

	fmt.Println("获取参数", r.GetString("gmt_create"))
	var noti, _ = utils.GetTradeNotification(r)
	if noti == nil {
		log.Alipay_log().Println("解析返回值异常", r.URL.String(), "data:", r.GetBodyString())
		//r.Response.Write("success") //确认收到通知
		return
	}
	fmt.Println("交易状态为:", noti.TradeStatus)

	//查询是否存在此账单
	status, err := Tripartite_api.Alipay_Order_payment_status(noti.OutTradeNo)
	if err != nil {
		log.Alipay_log().Println("支付结果查询账单失败", err.Error())
		return
	}
	if status == false {
		log.Alipay_log().Println("支付结果查询账单返回false")
		return
	}

	//支付成功逻辑....

	//获取用户订单
	userid, money, status_zhifu, err := Data.Data_get_number_dingdan(noti.OutTradeNo)
	if err != nil {
		log.Alipay_log().Println("支付宝支付成功回调，回调错误，未找到用户订单", err.Error(), "商户订单:", noti.OutTradeNo, "支付宝订单:", noti.TradeNo, "交易金额:", noti.TotalAmount)
		return
	}
	if status_zhifu == 1 {
		log.Alipay_log().Println("支付宝支付成功回调，回调错误，该订单数据库中已经为支付成功", "商户订单:", noti.OutTradeNo, "支付宝订单:", noti.TradeNo, "交易金额:", noti.TotalAmount)
		return
	}

	//添加用户余额
	Data.Data_Add_User_money(userid, money)
	//改变用户订单为支付成功
	Data.Data_Update_status(noti.OutTradeNo)

	r.Response.Write("success") //确认收到通知
}
