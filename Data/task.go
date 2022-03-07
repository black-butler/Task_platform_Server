package Data

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"platform/Bean"
	"platform/log"
	"platform/utils"
	"time"
)

//添加任务
func Data_add_task(user *Bean.User, body string, img string, one_money int, sum int) error {
	_, err := g.DB().Model("tasks").Data(g.Map{"userid": user.Id, "body": body, "imgs": img, "one_money": one_money, "sum": sum, "time": time.Now().Format(utils.Time_Format)}).Insert()
	if err != nil {
		log.Sql_log().Line().Println("添加任务", err.Error())
		return errors.New("添加任务失败")
	}
	return nil
}
