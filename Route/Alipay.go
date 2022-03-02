package Route

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
		r.Response.Write("success") //确认收到通知
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

	r.Response.Write("success") //确认收到通知
}
