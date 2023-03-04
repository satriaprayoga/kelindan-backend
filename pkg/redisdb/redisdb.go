package redisdb

import (
	"fmt"
	"kelindan/pkg/settings"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Setup() {
	now := time.Now()
	conString := fmt.Sprintf("%s:%s", settings.AppConfigSetting.RedisDB.Host, settings.AppConfigSetting.RedisDB.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     conString,
		Password: settings.AppConfigSetting.RedisDB.Password,
		DB:       settings.AppConfigSetting.RedisDB.DB,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	timeSpent := time.Since(now)
	log.Printf("Config redis is ready in %v", timeSpent)
}

// AddSession :
func AddSession(key string, val interface{}, mn time.Duration) error {
	set, err := rdb.Set(key, val, mn).Result()
	if err != nil {
		return err
	}
	fmt.Println(set)
	return nil
}

// GetSession :
func GetSession(key string) interface{} {
	value := rdb.Get(key).Val()
	fmt.Println(value)
	return value
}
