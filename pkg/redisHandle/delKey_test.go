package redishandle

import (
	"context"
	"fmt"
	"github.com/TremblingV5/DouTok/config/configStruct"
	"github.com/TremblingV5/DouTok/pkg/configurator"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"testing"
)

func TestKeys(t *testing.T) {
	redisConfig := configStruct.RedisConfig{}
	err := configurator.InitConfig(&redisConfig, "relation_redis.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	clientCache, err := InitRedis(redisConfig.Host+":"+redisConfig.Port, redisConfig.Password, redisConfig.Databases)
	if err != nil {
		fmt.Println(err)
		return
	}
	rd := clientCache[constants.DbDefault]
	res, err := rd.Client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	if err := rd.DelKey(context.Background(), "test_keyss"); err != nil {
		panic(err)
	}
}
