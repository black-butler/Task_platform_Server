package Data

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/log"
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
	_, err = g.DB().Model("users").Data(g.Map{"number": number, "password": password}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("添加用户", err.Error())
		return errors.New("账号已存在")
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
