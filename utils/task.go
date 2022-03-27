package utils

import (
	"errors"
	"github.com/gogf/gf/os/gcache"
	"platform/log"
	"sync"
	"time"
)

//获取接任务锁
var Jie_Task_Suo *gcache.Cache
var Cache_mutux sync.Mutex

func init() {
	Jie_Task_Suo = gcache.New()
}

//获取某个任务的锁
func Get_task_suo(taskid int) (*sync.Mutex, error) {
	Cache_mutux.Lock()
	defer Cache_mutux.Unlock()

	v, err := Jie_Task_Suo.Get(taskid)
	if err != nil {
		log.File_gcache_log().Println("获取某个任务的锁错误", err.Error())
		return nil, err
	}
	if v == nil {
		Jie_Task_Suo.Set(taskid, new(sync.Mutex), time.Minute*120)
	}

	v, err = Jie_Task_Suo.Get(taskid)
	if err != nil {
		log.File_gcache_log().Println("获取某个任务的锁错误", err.Error())
		return nil, err
	}
	zhi, ok := v.(*sync.Mutex)
	if ok == false {
		return nil, errors.New("获取任务锁错误")
	}

	return zhi, nil
}
