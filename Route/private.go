package Route

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	g := g.Server()
	group := g.Group("/private")

	//公告接口
	group.GET("/announcement", announcement)
	//单子列表
	group.GET("/order_list", order_list)
	//单子详情页
	group.GET("/detail", detail)
	//发单子
	group.GET("/submit", submit)
	//上传图片
	group.GET("/uploading_img", uploading)
	//用户充值
	group.GET("/user_top_up", user_top_up)
}

//公告接口
func announcement(r *ghttp.Request) {

}

//单子列表
func order_list(r *ghttp.Request) {

}

//单子详情页接口
func detail(r *ghttp.Request) {

}

//提交单子接口
func submit(r *ghttp.Request) {

}

//上传图片
func uploading(r *ghttp.Request) {

}

//用户充值
func user_top_up(r *ghttp.Request) {}
