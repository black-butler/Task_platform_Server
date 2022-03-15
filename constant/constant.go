package constant

//任务状态
const (
	Zhengchang = 0 //正常
	Xiajia     = 1 //下架
)

var Task_status_map = map[int]string{Zhengchang: "正常", Xiajia: "下架"}

//接单状态
const (
	Weiwancheng = 0 //未完成
	Yiwancheng  = 1 //已完成并打款
	Chaoshi     = 3 //超时
	Word_XiaJIa = 4 //下架
)
