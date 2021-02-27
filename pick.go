package redis

import (
	"context"
	"reflect"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

func Pick(id string, tc interface{}) error {

	// 反射值和类型
	tcValue := reflect.ValueOf(tc).Elem()
	tcType := reflect.TypeOf(tc).Elem()
	// 读取配置文件
	redisConf, err := ini.Load("conf/app.ini")
	if err != nil {
		return errors.Wrap(err, "loading conf file err when pick token from redis")
	}
	// 读取redis配置
	redisAddr := redisConf.Section("redis").Key("addr").String()
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer redisClient.Close()
	// 循环读取token相关数据
	for i := 0; i < tcValue.NumField(); i++ {
		v, err := redisClient.HGet(context.Background(), id, tcType.Field(i).Name).Result()
		if err != nil {
			return errors.Wrap(err, "pick token from redis error")
		}
		// 写入token相关数据
		tcValue.Field(i).SetString(v)
	}
	// 返回结果
	return nil
}
