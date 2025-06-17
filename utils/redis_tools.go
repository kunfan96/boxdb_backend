package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb *redis.Client

func InitRedis() {
	var RedisCfg = GetBootstrapConfig().Redis

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(RedisCfg.Host, ":", RedisCfg.Port),
		Password: RedisCfg.Password,
		DB:       0,
	})

	// test connect
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		Logger.Error(fmt.Sprintf("Redis Connect Error:%s", err))
	} else {
		Logger.Info(fmt.Sprintf("Redis Connect Success:%s", pong))
	}
}

// set redis hash
func SetRedisStringByKey(key, value string) {
	rdb.Set(ctx, key, value, 48*time.Hour) // 设置整个 String 2 天过期
}

// get redis hash value
func GetRedisStringByKey(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// delete redis string
// return:
//   - 0 delete happen error
//   - 1 delete success
//   - 2 no such key
func DelRedisStringByKey(key string) int8 {
	exists, _ := rdb.Exists(ctx, key).Result()

	if exists > 0 {
		err := rdb.Del(ctx, key).Err()

		if err != nil {
			Logger.Error(fmt.Sprintf("Del Redis key %s is Error:%s", key, err))

			return 0
		}

		return 1
	}

	return 2

}
