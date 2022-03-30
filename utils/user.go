package utils

import (
	"errors"
	"github.com/gogf/gf/os/gcache"
	"platform/log"
	"sync"
	"time"
)

//获取接任务锁
var User_Suo *gcache.Cache
var Cache_User_mutux sync.Mutex

func init() {
	User_Suo = gcache.New()
}

//获取某个用户的锁
func Get_user_suo(userid int) (*sync.Mutex, error) {
	Cache_User_mutux.Lock()
	defer Cache_User_mutux.Unlock()

	v, err := User_Suo.Get(userid)
	if err != nil {
		log.File_gcache_log().Println("获取某个用户的锁错误", err.Error())
		return nil, err
	}
	if v == nil {
		User_Suo.Set(userid, new(sync.Mutex), time.Minute*120)
	}

	v, err = User_Suo.Get(userid)
	if err != nil {
		log.File_gcache_log().Println("获取某个用户的锁错误", err.Error())
		return nil, err
	}
	zhi, ok := v.(*sync.Mutex)
	if ok == false {
		return nil, errors.New("获取用户锁错误")
	}

	return zhi, nil
}
