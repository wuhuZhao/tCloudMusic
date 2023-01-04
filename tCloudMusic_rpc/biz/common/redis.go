package common

import (
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/go-redis/redis"
)

type Redis struct {
	RedisHost string
	RedisPort uint16
	RedisAuth string
	Database  int
	db        *redis.Client
}

// Init: init redis connection
func (r *Redis) Init() (err error) {
	redisAddr := fmt.Sprintf("%s:%d", r.RedisHost, r.RedisPort)
	r.db = redis.NewClient(&redis.Options{
		Addr:        redisAddr,   // redis addr
		Password:    r.RedisAuth, // redis password
		DB:          r.Database,  // use database
		IdleTimeout: 300,         // timeout
		PoolSize:    100,         // connect pool
	})
	klog.Infof("connecting redis : %v\n", redisAddr)
	_, err = r.db.Ping().Result()
	if err != nil {
		klog.Errorf("contect redis failed! err: %v\n", err)
		return err
	}
	klog.Infof("connect redis success!")
	return nil
}

// Insert: redis command set : input (key, value, expire) to redis
func (r *Redis) Insert(key string, value interface{}, expire time.Duration) (err error) {
	_, err = r.db.Set(key, value, expire).Result()
	return
}

// Delete: redis command del : input key to delete redis key
func (r *Redis) Delete(key string) (err error) {
	_, err = r.db.Del(key).Result()
	return
}
