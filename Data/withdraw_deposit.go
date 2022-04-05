package Data

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"platform/constant"
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

//申请提现列表
func Data_withdraw_deposit_apply_for() (gdb.Result, error) {
	result, err := g.DB().Model("withdraw_deposit").Where("status", constant.TiXian_weiwancheng).All()
	if err != nil {
		log.Sql_log().Println("申请提现列表错误" + err.Error())
		return nil, err
	}
	return result, nil
}

//获取某个id的提现记录
//func Data_id_withdraw_deposit(id int)(gdb.Record,error){
//
//	record,err :=
//
//
//}

//更改某个id的提现状态
func Data_update_deposit_apply_status(id int, status int) error {
	_, err := g.DB().Model("withdraw_deposit").Data(g.Map{"status": status}).Where("id", id).Update()
	if err != nil {
		log.Sql_log().Println("更改某个id的提现状态错误" + err.Error())
		return err
	}
	return nil
}
