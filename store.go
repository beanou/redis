// redis 存储功能
package redis

import (
	"context"
	"reflect"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

func Store(id string, tc interface{}, expire time.Duration) error {

	// 反射值和类型
	tcValue := reflect.ValueOf(tc).Elem()
	tcType := reflect.TypeOf(tc).Elem()
	// 读取配置文件
	redisConf, err := ini.Load("conf/app.ini")
	if err != nil {
		return errors.Wrap(err, "error to read conf file")
	}
	// 读取redis配置
	redisAddr := redisConf.Section("redis").Key("addr").String()
	redisClient := redis.NewClient(&redis.Options{Addr: redisAddr})
	defer redisClient.Close()
	// 哈希方式写入token相关属性到redis
	for i := 0; i < tcValue.NumField(); i++ {
		err = redisClient.HSet(context.TODO(), id, tcType.Field(i).Name, tcValue.Field(i).String()).Err()
		if err != nil {
			return errors.Wrap(err, "redis store err")
		}
	}
	// 设置过期时间
	err = redisClient.Expire(context.Background(), id, time.Second*expire).Err()
	if err != nil {
		return errors.Wrap(err, "redis set expire err")
	}
	// 返回结果
	return nil
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
