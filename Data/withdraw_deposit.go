package Data

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"platform/log"
	"platform/utils"
	"time"
)

//添加提现记录
func Data_Add_withdraw_deposit(userid int, money int, alipay_number string, alipay_name string) error {
	_, err := g.DB().Model("withdraw_deposit").Data(g.Map{
		"userid":        userid,
		"money":         money,
		"time":          time.Now().Format(utils.Time_Format),
		"alipay_number": alipay_number,
		"alipay_name":   alipay_name,
		"status":        0}).
		Insert()
	if err != nil {
		log.Sql_log().Println("添加提现记录错误" + err.Error())
		return err
	}
	return nil
}

//提现记录
func Data_withdraw_deposit_record(userid int) (gdb.Result, error) {
	result, err := g.DB().Model("withdraw_deposit").Where("userid", userid).All()
	if err != nil {
		log.Sql_log().Println("提现记录错误" + err.Error())
		return nil, err
	}
	return result, nil
}
