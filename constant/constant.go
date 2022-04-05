package constant

//任务状态
const (
	Zhengchang = 0 //正常
	Xiajia     = 1 //下架
	Weishenhe  = 2 //未审核
	Butongguo  = 3 //审核不通过
)

var Task_status_map = map[int]string{Zhengchang: "正常", Xiajia: "下架"}

//工单状态
const (
	Weiwancheng = 0 //未完成
	Yiwancheng  = 1 //已完成并打款
	Chaoshi     = 3 //超时
	Word_XiaJIa = 4 //下架
)

//提现状态
const (
	TiXian_weiwancheng = 0 //提现未完成
	Tixian_yiwancheng  = 1 //提现已完成
)
