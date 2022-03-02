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
