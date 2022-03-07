package main

import (
	"github.com/gogf/gf/frame/g"
	"os"
	"path/filepath"

	_ "platform/Route" //初始化路由
)

func main() {
	//获取程序绝对路径
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//设置系统日志
	g.Log().SetPath(dir + "\\glog")

	g := g.Server()
	//g.SetHTTPSPort(9090)
	//g.EnableHTTPS("config/server.crt", "config/server.key")
	g.SetPort(9090)
	g.Run()
}
