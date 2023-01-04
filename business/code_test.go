package business

import (
	"contact/cache"
	"contact/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func initRedis() {
	global.Redis = *redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
		Username: "",
		Password: "",
		DB:       0,
	})
}

//func Test TestUserCodeBusiness_RandomCodes(t *testing.T) {
func TestGroupCodeBusiness_RandomCode(t *testing.T) {
	initRedis()
	b := GroupCodeBusiness{}
	codes, err := b.RandomCodes(11, false)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	for _, c := range codes {
		fmt.Printf("get userCode: %d\n", c)
	}

	// 清除缓存(需要生成的话不做清除即可)
	r := cache.RedisServer{}
	r.FuzzyClear("qvbilam:*")
}

func TestClear(t *testing.T) {
	initRedis()
	r := cache.RedisServer{}
	r.FuzzyClear("qvbilam:*")
}

func TestSetNx(t *testing.T) {
	initRedis()
	r := cache.RedisServer{}
	res := r.SetNX("qvbilam:tmd", "123", 0)
	fmt.Println("结果: ", res)
}
