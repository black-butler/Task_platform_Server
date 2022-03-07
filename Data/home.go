package Data

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"platform/log"
)

//获取首页数据
func Data_Get_home() (string, string, error) {
	result, err := g.DB().Model("home_data").One()
	if err != nil {
		log.Sql_log().Line().Println("获取首页数据失败", err.Error())
		return "", "", errors.New("读取文件名失败")
	}

	if len(result) <= 0 {
		return "", "", errors.New("获取首页数据失败")
	}

	return result["lunbo_imgs"].String(), result["gonggao"].String(), nil
}
