// redis 存储功能
package redis

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
)

func Stroe(id string, tc *TokenContent, expire time.Duration) Results {

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
		er := redisClient.HSet(context.TODO(), id, tcType.Field(i).Name, tcValue.Field(i).String()).Err()
		fmt.Println(er)
	}
	redisClient.Expire(context.Background(), id, time.Second*expire)

	// rs.Data = tc
	rs.ErrCode = 0
	rs.ErrMsg = "ok"

	return rs
}

// func StroeAll(id string, tc *TokenContent, expire int64) Results {
//
// var rs Results
// tcValue := reflect.ValueOf(tc).Elem()
// tcType := reflect.TypeOf(tc).Elem()
// fmt.Println(tcValue.NumField(), tcType)
//
// redisConf, err := ini.Load("conf/app.ini")
// if err != nil {
//   fmt.Println(err)
// }
//
// redisAddr := redisConf.Section("redis").Key("addr").String()
// redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
// defer redisClient.Close()
//
// m := make(map[string]interface{})
//
// for i := 0; i < tcValue.NumField(); i++ {
//   fmt.Println(tcType.Field(i).Name, "----", tcValue.Field(i))
//   m[tcType.Field(i).Name] = tcValue.Field(i).String()
//   fmt.Println(m)
//   // er := redisClient.HSet(context.TODO(), id, tcType.Field(i).Name, tcValue.Field(i).String()).Err()
// }
//
// er := redisClient.HMSet(context.TODO(), id, m).Err()
// // er := redisClient.HMSet(context.TODO(), id, tc).Err() // first ,need  func (tc *TokenContent) MarshalBinary() ([]byte, error) //and then  it has an err: ERR wrong number of arguments for 'hmset' command
// fmt.Println(er)
//
// rs.Data = tc
// rs.ErrCode = 0
// rs.ErrMsg = "ok"
//
// return rs
// }
