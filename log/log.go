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

func File_read() *glog.Logger {
	return g.Log("File_read").File("文件读取日志-{Ymd}.log")
}

func File_gcache_log() *glog.Logger {
	return g.Log("File_gcache_log").File("gcache缓存日志-{Ymd}.log")
}

func File_timed_log() *glog.Logger {
	return g.Log("File_timed").File("定时任务日志-{Ymd}.log")
}

func File_core_log() *glog.Logger {
	return g.Log("File_core_log").File("业务核心逻辑日志-{Ymd}.log").Line()
}

func File_admin_log() *glog.Logger {
	return g.Log("File_admin_log").File("管理员操作日志-{Ymd}.log").Line()
}
