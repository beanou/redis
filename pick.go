package redis

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

func Pick(id string, tc *TokenContent) Results {
	var rs Results

	tcValue := reflect.ValueOf(tc).Elem()
	tcType := reflect.TypeOf(tc).Elem()
	fmt.Println(tcValue.NumField(), tcType)

	redisConf, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Println(err)
	}

	redisAddr := redisConf.Section("redis").Key("addr").String()
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer redisClient.Close()

	for i := 0; i < tcValue.NumField(); i++ {
		fmt.Println(tcType.Field(i).Name, "----", tcValue.Field(i))
		v, err := redisClient.HGet(context.Background(), id, tcType.Field(i).Name).Result()
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println(v)
		tcValue.Field(i).SetString(v)
	}

	// rs.Data = tc
	rs.ErrCode = 0
	rs.ErrMsg = "ok"
	return rs
}
