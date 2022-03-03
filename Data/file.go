package Data

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/log"
	"platform/utils"
	"time"
)

//保存文件
func Data_Save_file(user *Bean.User, filename string) (int, error) {
	id, err := g.DB().Model("imgs").Data(g.Map{"filename": filename, "userid": user.Id, "time": time.Now().Format(utils.Time_Format)}).InsertAndGetId()
	if err != nil {
		log.Sql_log().Line().Println("保存文件地址失败", err.Error())
		return 0, errors.New("文件保存失败")
	}
	return int(id), nil
}

//通过id获取图片名
func Get_Img_filename(id int) (string, error) {

	result, err := g.DB().Model("imgs").Data(g.Map{"id": id}).One()
	if err != nil {
		log.Sql_log().Line().Println("读取文件名失败", err.Error())
		return "", errors.New("读取文件名失败")
	}

	if len(result) <= 0 {
		return "", errors.New("找不到此文件")
	}

	if filename, ok := result["filename"]; ok {
		return filename.Val().(string), nil
	} else {
		return "", errors.New("找不到此文件")
	}
}

//查找是否存在此文件
func Check_fileid(fileid int) error {
	result, err := g.DB().Model("imgs").Where("id", fileid).One()
	if err != nil {
		log.Sql_log().Line().Println("查找是否存在此文件 查找失败", err.Error())
		return errors.New("文件不存在")
	}
	if len(result) == 0 {
		log.Sql_log().Line().Println("查找是否存在此文件 查找失败")
		return errors.New("文件不存在")
	}

	return nil
}
