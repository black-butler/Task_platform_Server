package Tripartite_api

import (
	"errors"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"platform/utils"
)

//查询支付宝是否存在此订单并且是否支付
func Alipay_Order_payment_status(commercial_tenant_id string) (bool, error) {

	var p = alipay.TradeQuery{}
	p.OutTradeNo = commercial_tenant_id //商户订单号

	Zhi, err := utils.Client.TradeQuery(p)
	if err != nil {
		fmt.Println("pay client.TradeAppPay error:", err)
		return false, errors.New("验证失败")
	}

	if Zhi.Content.TradeStatus == "TRADE_SUCCESS" || Zhi.Content.TradeStatus == "TRADE_FINISHED" {
		//交易支付成功
		return true, nil
	} else {
		return false, nil
	}
}
