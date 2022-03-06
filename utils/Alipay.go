package utils

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/smartwalle/alipay/v3"
)

var Client *alipay.Client

func init() {
	// appId
	appId := "2021002194600072"
	// 应用公钥
	//aliPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAigqcp8ouSzcykTpCvw6rDkrA6UUAnjstIzb/sTmy5Qnsbyyk6kasJUshmm0ZXMmWeOkERhwMtaCgwVEb/eXEk7ibrCWLyiBFFggFzjkLvHv2AzjYpbL8mum788cEPANxgm5O8rSSqrUrXVTl/tUtVBLV0QByy3UNvGlcUd931DDIco1FICbyAGli6TmaqIwgYPU3huAvNXww/zgDNJQ5iU+3fi2afPlsNDA96LCkRQvept7QqlDE70fe5L9H9e/0iCE38US6Y75rUgNuClcdq79gjloJGSFQfG2GBkDT6SY1E+kh2tdaPA2RztfhHC69xojrJNSeBNpNkWpktkNphwIDAQAB"
	// 应用私钥
	privateKey := "MIIEowIBAAKCAQEAigqcp8ouSzcykTpCvw6rDkrA6UUAnjstIzb/sTmy5Qnsbyyk6kasJUshmm0ZXMmWeOkERhwMtaCgwVEb/eXEk7ibrCWLyiBFFggFzjkLvHv2AzjYpbL8mum788cEPANxgm5O8rSSqrUrXVTl/tUtVBLV0QByy3UNvGlcUd931DDIco1FICbyAGli6TmaqIwgYPU3huAvNXww/zgDNJQ5iU+3fi2afPlsNDA96LCkRQvept7QqlDE70fe5L9H9e/0iCE38US6Y75rUgNuClcdq79gjloJGSFQfG2GBkDT6SY1E+kh2tdaPA2RztfhHC69xojrJNSeBNpNkWpktkNphwIDAQABAoIBAHUJAkCP/ifvKIRQrP5nZUe8wUoFIr0E1wQgQTD1BmOBwl+PrlKikJkd1eOj0/kDQPxKM0Ftzqi/AwxjxCPDhqLnxbRyMM6yBWDkdGefnf+z6aRZsfTqh5ifyqaoeUKYeMho1a6YwjDRYW7D6CvieYfqYDXO94TgPUEc4xTXLjVL0OS1YgCAx6eZoWKJWz1v/Xq+jo3WUauLc+6F928iUGcnNIENg/y0tbymuTbZ8cr9ND2sldnsHLd3p2rUYUjQfJkTWXXLsDJFyBClpAGkObcqVCbpwFedrtT9s1lP+KvQCQ8W3j2/PbltwaTxReJeVT8tZ+RnEuxdaTbud9pN44ECgYEA6JAnDRFm0Er12nhHP3i4KrN708ewCuRBe1rFvjXMQ87MVNfioi7rV5buExx+1YBDOur4X32dq/Lp4D4Digu1wxXNWyiFkF6v6Rq9Y8aSQjzGCXOBlFes+MgdHb8PpGylJNWdlX+dkJvG5Ca2P7h5D5rV3Qj4YdUs6CwO7uNSq3ECgYEAl/PqO7diH6SZ4S1RPW3h89J2HybLATN6vCg+RY+8s6czg8i6S8KwsELTAxgI2uQLVqGifu3+kknqYAUOEIcNOCqCXbWmF1xBKOWWaCV4ZbqgKh0MWytuc9CjJyKrKMK8DSIBskMrNmkZ+S+EapUsM0cg15ZA3JQBHNauDA4lOHcCgYAYnOkFIQpYkRZkAMbJmOUk38oDJ+chv/aOL5UuBFOR+Zj2gcKil4SgyIB51VI3FlQHMEcJFCpTwGmwKeAGBCdAdlY9h5RbKypC6WmR3bos+HGdHnRgVscfrU4nj8kABd+UfmcnI1Jxs4rhKpevNr7ZP/HSatiewgj2qXMLJVPigQKBgFxuuZud0Aijnh+F65dMklg5PDVy6aZPZGe0qzyxVP6LxSBzKDARvF1cKPQG2MweUG9gX3KK34Kph/Lk4EtZe8cgxLCwYNpw+gogrr+nm3d2cRttFCkZYFT/I2AZDLj8zFvIxfNkPJMal/wm1YvoNjzzFZ1O/yGuvoaGaNVYfXe9AoGBALoBMFqfk7uq/s/mi0KGs8rOt0J12F+TOhGiSN2z5ExPzqZYIvKEAjLIJaTA6p89MVwa0YCzMqEzMRY3SCdS4pisYEfeDXMJzz4ONmwnk1y+HX0e6rnBqy+S/iSJqL191MOpbgfDYq5Jx9ISIMc0nc0t7KtCyOJcGVIuZ1jjE2Zb"
	var err error
	Client, err = alipay.New(appId, privateKey, true)
	if err != nil {
		panic(err)
	}

	Client.LoadAppPublicCertFromFile("config/appCertPublicKey_2021002194600072.crt") // 加载应用公钥证书
	Client.LoadAliPayRootCertFromFile("config/alipayRootCert.crt")                   // 加载支付宝根证书
	Client.LoadAliPayPublicCertFromFile("config/alipayCertPublicKey_RSA2.crt")       // 加载支付宝公钥证书
}

//支付宝验签
func GetTradeNotification(req *ghttp.Request) (noti *alipay.TradeNotification, err error) {

	noti = &alipay.TradeNotification{}
	noti.AppId = req.GetString("app_id")
	noti.AuthAppId = req.GetString("auth_app_id")
	noti.NotifyId = req.GetString("notify_id")
	noti.NotifyType = req.GetString("notify_type")
	noti.NotifyTime = req.GetString("notify_time")
	noti.TradeNo = req.GetString("trade_no")
	noti.TradeStatus = alipay.TradeStatus(req.GetString("trade_status"))
	noti.TotalAmount = req.GetString("total_amount")
	noti.ReceiptAmount = req.GetString("receipt_amount")
	noti.InvoiceAmount = req.GetString("invoice_amount")
	noti.BuyerPayAmount = req.GetString("buyer_pay_amount")
	noti.SellerId = req.GetString("seller_id")
	noti.SellerEmail = req.GetString("seller_email")
	noti.BuyerId = req.GetString("buyer_id")
	noti.BuyerLogonId = req.GetString("buyer_logon_id")
	noti.FundBillList = req.GetString("fund_bill_list")
	noti.Charset = req.GetString("charset")
	noti.PointAmount = req.GetString("point_amount")
	noti.OutTradeNo = req.GetString("out_trade_no")
	noti.OutBizNo = req.GetString("out_biz_no")
	noti.GmtCreate = req.GetString("gmt_create")
	noti.GmtPayment = req.GetString("gmt_payment")
	noti.GmtRefund = req.GetString("gmt_refund")
	noti.GmtClose = req.GetString("gmt_close")
	noti.Subject = req.GetString("subject")
	noti.Body = req.GetString("body")
	noti.RefundFee = req.GetString("refund_fee")
	noti.Version = req.GetString("version")
	noti.SignType = req.GetString("sign_type")
	noti.Sign = req.GetString("sign")
	noti.PassbackParams = req.GetString("passback_params")
	noti.VoucherDetailList = req.GetString("voucher_detail_list")
	noti.AgreementNo = req.GetString("agreement_no")
	noti.ExternalAgreementNo = req.GetString("external_agreement_no")

	//if len(noti.NotifyId) == 0 {
	//	return nil, errors.New("不是有效的 Notify")
	//}
	//req.ParseForm(req.Form)
	ok, err := Client.VerifySign(req.Request.Form)
	if ok == false {
		return nil, err
	}
	return noti, err
}
