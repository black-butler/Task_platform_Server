package Data

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/log"
	"platform/utils"
	"time"
)

//创建订单
func Data_pay_Create_dingdan(user *Bean.User, order_number string, money int) error {
	_, err := g.DB().Model("order_form").Data(g.Map{"order_number": order_number, "userid": user.Id, "money": money, "status": 0, "time": time.Now().Format(utils.Time_Format)}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("支付宝创建订单失败", err.Error())
		return errors.New("创建订单失败")
	}
	return nil
}

//根据订单号获取订单
func Data_get_number_dingdan(order_number string) (string, int, error) {
	result, err := g.DB().Model("order_form").Where("order_number", order_number).One()
	if err != nil {
		log.Sql_log().Line().Println("根据订单号获取订单失败", err.Error())
		return "", 0, errors.New("根据订单号获取订单")
	}

	if len(result) <= 0 {
		log.Sql_log().Line().Println("根据订单号获取订单失败，订单号为空，查询的订单号:", order_number)
		return "", 0, errors.New("根据订单号获取订单")
	}

	zhi_userid, ok := result["userid"]
	if ok == false {
		log.Sql_log().Line().Println("根据订单号获取订单失败，订单号为空，查询的订单号:", order_number)
		return "", 0, errors.New("根据订单号获取订单")
	}

	zhi_money, ok := result["money"]
	if ok == false {
		log.Sql_log().Line().Println("根据订单号获取订单失败，订单号为空，查询的订单号:", order_number)
		return "", 0, errors.New("根据订单号获取订单")
	}

	return zhi_userid.String(), zhi_money.Int(), nil
}

//改变订单成功状态
func Data_Update_status(order_number string) {
	_, err := g.DB().Model("order_form").Data(g.Map{"status": 1}).Where("order_number", order_number).Update()
	if err != nil {
		log.Sql_log().Line().Println("改变订单成功状态失败", err.Error())
		return
	}
}
