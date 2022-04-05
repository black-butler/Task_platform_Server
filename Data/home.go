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

//更换首页轮播图
func Data_update_home_lun_img(zhi string) error {
	_, err := g.DB().Model("home_data").Data(g.Map{"lunbo_imgs": zhi}).Where("id", "1").Update()
	if err != nil {
		log.Sql_log().Line().Println("更换首页轮播图失败", err.Error())
		return errors.New("更换首页轮播图失败")
	}
	return nil
}
