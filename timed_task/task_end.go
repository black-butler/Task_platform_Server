package timed_task

import (
	"platform/Data"
	"platform/constant"
	"platform/log"
	"platform/utils"
	"time"
)

//检测任务是否超时
func init() {
	go func() {
		log.File_timed_log().Println("检测任务是否超时 - 定时任务开启成功")
		for {

			tasks, err := Data.Data_get_all_task()
			if err != nil {
				time.Sleep(time.Minute)
				continue
			}
			for _, v := range tasks {
				//判断该任务是否已经下架
				ti, err := time.Parse(utils.Time_Format, v["endDate"].String())
				if err != nil {
					log.File_timed_log().Println("格式化时间错误任务id", v["id"].Int(), "结束时间:", v["endDate"].String())
					continue
				}

				if time.Now().Unix() > ti.Unix() {
					//更新任务状态
					Data.Data_update_task_status(v["id"].Int(), constant.Xiajia)
				}
			}

			time.Sleep(30 * time.Minute)
		}
	}()
}
