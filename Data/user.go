package Data

import (
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/log"
	"strconv"
	"sync"
)

var mu sync.Mutex

//添加用户
func Data_Add_user(number string, password string) error {
	mu.Lock()
	defer mu.Unlock()

	result, err := g.DB().Model("users").Where("number", number).One()
	if err != nil {
		log.Sql_log().Line().Println("数据库添加用户 查询用户失败", err.Error())
		return errors.New("账号已存在")
	}
	if len(result) != 0 {
		log.Sql_log().Line().Println("数据库添加用户 用户已存在")
		return errors.New("账号已存在")
	}
	//println(result)
	_, err = g.DB().Model("users").Data(g.Map{"number": number, "password": password, "img": 1, "alipay_number": "", "alipay_name": ""}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("添加用户", err.Error())
		return errors.New("账号已存在")
	}
	return nil
}

//更新用户支付宝账户和名字
func Data_update_user_alipay(userid int, alipay_number string, alipay_name string) error {
	_, err := g.DB().Model("users").Data(g.Map{"alipay_number": alipay_number, "alipay_name": alipay_name}).Where("id", userid).Update()
	if err != nil {
		log.Sql_log().Println("更新用户支付宝账户和名字错误:", err.Error())
		return err
	}
	return nil
}

//查找用户 通过用户账号密码
func Data_Get_user(number string, password string) (*Bean.User, error) {
	user := new(Bean.User)
	err := g.DB().Model("users").Where("number", number).Where("password", password).Struct(user)
	if err != nil {
		log.Sql_log().Line().Println("查找用户", err.Error())
		return nil, errors.New("账号不存在")
	}
	return user, nil
}

//查找用户 通过用户id
func Data_Get_userid(userid string) (*Bean.User, error) {
	user := new(Bean.User)
	err := g.DB().Model("users").Where("id", userid).Struct(user)
	if err != nil {
		log.Sql_log().Line().Println("查找用户", err.Error())
		return nil, errors.New("账号不存在")
	}
	return user, nil
}

//刷新用户
func Data_refre_userid(user *Bean.User) error {
	err := g.DB().Model("users").Where("id", user.Id).Struct(user)
	if err != nil {
		log.Sql_log().Line().Println("查找用户", err.Error())
		return errors.New("账号不存在")
	}
	return nil
}

//更新用户头像id
func Data_Update_User_touxiangid(user *Bean.User, touxiangid int) error {
	_, err := g.DB().Model("users").Data(g.Map{"img": touxiangid}).Where("id", user.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("更新用户头像失败", err.Error())
		return errors.New("更新用户头像失败")
	}

	return nil
}

//添加用户余额
func Data_Add_User_money(userid string, money int) {
	_, err := g.DB().Model("users").Data(g.Map{"money": gdb.Raw("money+" + strconv.Itoa(money))}).Where("id", userid).Update()
	if err != nil {
		log.Sql_log().Line().Println("添加用户余额失败", err.Error())
		return
	}
}

//扣除用户余额
func Data_delete_user_money(userid int, money int) error {
	_, err := g.DB().Model("users").Data(g.Map{"money": gdb.Raw("money-" + strconv.Itoa(money))}).Where("id", userid).Update()
	if err != nil {
		log.Sql_log().Line().Println("添加用户余额失败", err.Error())
		return err
	}
	return nil
}

//扣除用户余额到冻结余额
func Data_transfer_money(user *Bean.User, money int) error {
	_, err := g.DB().Model("users").Data(g.Map{"money": gdb.Raw("money-" + strconv.Itoa(money)), "freeze_money": gdb.Raw("freeze_money+" + strconv.Itoa(money))}).Where("id", user.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("扣除用户余额到冻结余额失败", err.Error())
		return errors.New("用户余额操作失败")
	}

	return nil
}

//扣除用户冻结余额到余额
func Data_transfer_money_freeze(user *Bean.User, money int) error {
	_, err := g.DB().Model("users").Data(g.Map{"money": gdb.Raw("money+" + strconv.Itoa(money)), "freeze_money": gdb.Raw("freeze_money-" + strconv.Itoa(money))}).Where("id", user.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("扣除用户余额到冻结余额失败", err.Error())
		return errors.New("用户余额操作失败")
	}
	return nil
}

//扣除用户冻结余额
func Data_delete_user_freeze_money(user *Bean.User, money int) error {
	_, err := g.DB().Model("users").Data(g.Map{"freeze_money": gdb.Raw("freeze_money-" + strconv.Itoa(money))}).Where("id", user.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("扣除用户冻结余额失败", err.Error())
		return errors.New("用户余额操作失败")
	}
	return nil
}

//添加用户正常余额
func Data_add_user_money(user *Bean.User, money int) error {
	_, err := g.DB().Model("users").Data(g.Map{"money": gdb.Raw("money+" + strconv.Itoa(money))}).Where("id", user.Id).Update()
	if err != nil {
		log.Sql_log().Line().Println("添加用户正常余额", err.Error())
		return errors.New("用户余额操作失败")
	}
	return nil
}
