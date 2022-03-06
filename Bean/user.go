package Bean

type User struct {
	Id            int
	Number        string
	Password      string
	Img           int    //头像文件id
	Money         int    //用户余额
	Freeze_money  int    //冻结金额
	Alipay_number string //绑定的支付宝账户
}
