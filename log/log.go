package log

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func Alipay_log() *glog.Logger {
	return g.Log("alipay").File("支付宝支付日志-{Ymd}.log")
}

func Sql_log() *glog.Logger {
	return g.Log("sql").File("数据库错误日志-{Ymd}.log")
}

func Login_log() *glog.Logger {
	return g.Log("Login_log").File("注册登录接口日志-{Ymd}.log")
}

func UploadFile_log() *glog.Logger {
	return g.Log("UploadFile_log").File("文件上传日志-{Ymd}.log")
}
