package cache

import (
	"contact/global"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

type RedisServer struct {
}

const DefaultGroupCodeDigit = 5

func (s *RedisServer) Exists(key string) int64 {
	result, _ := global.Redis.Exists(context.Background(), key).Result()
	return result
}

func (s *RedisServer) SetNX(key string, value string, expire int) bool {
	result, _ := global.Redis.SetNX(context.Background(), key, value, time.Duration(expire)*time.Second).Result()
	return result
}

func (s *RedisServer) Get(key string) string {
	result, _ := global.Redis.Get(context.Background(), key).Result()
	return result
}

func (s *RedisServer) Delete(key string, value string) int64 {
	if value != "" {
		result, _ := global.Redis.Get(context.Background(), key).Result()
		if result != value {
			return 0
		}
	}
	result, _ := global.Redis.Del(context.Background(), key).Result()
	return result
}

func (s *RedisServer) Clear(keys ...string) int64 {
	result, _ := global.Redis.Del(context.Background(), keys...).Result()
	return result
}

func (s *RedisServer) FuzzyClear(key string) int64 {
	keys := s.keys(key)
	if len(keys) == 0 {
		return 0
	}
	return s.Clear(keys...)
}

func (s *RedisServer) keys(key string) []string {
	result, _ := global.Redis.Keys(context.Background(), key).Result()
	return result
}

func (s *RedisServer) GetUserCodeDigit() int64 {
	key := RedisKey{}.GetGeneratorGroupCodeMaxDigit()
	value := s.Get(key)
	i, _ := strconv.Atoi(value)
	if i == 0 {
		return DefaultGroupCodeDigit
	}
	return int64(i)
}

func (s *RedisServer) GetGroupSpecialCodeDigit() int64 {
	key := RedisKey{}.GetGeneratorGroupSpecialCodeMaxDigit()
	value := s.Get(key)
	i, _ := strconv.Atoi(value)
	if i == 0 {
		return DefaultGroupCodeDigit
	}
	return int64(i)
}

func (s *RedisServer) SetGroupCodeDigit(digit int64) (string, error) {
	key := RedisKey{}.GetGeneratorGroupCodeMaxDigit()
	maxDigit := s.GetUserCodeDigit()
	if maxDigit >= digit {
		return "", status.Errorf(codes.InvalidArgument, "无法设置指定位数, 当前最大位数: %d", maxDigit)
	}
	val := strconv.Itoa(int(digit))
	result, _ := global.Redis.Set(context.Background(), key, val, 0).Result()

	return result, nil
}

func (s *RedisServer) SetGroupSpecialCodeDigit(digit int64) (string, error) {
	key := RedisKey{}.GetGeneratorGroupSpecialCodeMaxDigit()
	maxDigit := s.GetGroupSpecialCodeDigit()
	if maxDigit >= digit {
		return "", status.Errorf(codes.InvalidArgument, "无法设置指定位数, 当前最大位数: %d", maxDigit)
	}
	val := strconv.Itoa(int(digit))
	result, _ := global.Redis.Set(context.Background(), key, val, 0).Result()
	return result, nil
}

func (s *RedisServer) GenerateGroupCodes(digit int64, data []interface{}) (int64, error) {
	// 验证位数
	if _, err := s.SetGroupCodeDigit(digit); err != nil {
		return 0, err
	}

	// 设置锁
	lockKey := RedisKey{}.GetGeneratorGroupCodeLockKey(digit)
	if res := s.SetNX(lockKey, strconv.Itoa(int(digit)), 0); res == false {
		return 0, status.Errorf(codes.AlreadyExists, "当前位数已存在，key:%s", lockKey)
	}

	// 生成集合
	key := RedisKey{}.GetGroupCodesKey()
	result, _ := global.Redis.SAdd(context.Background(), key, data).Result()
	if result <= 0 {
		_ = s.Delete(lockKey, strconv.Itoa(int(digit))) // 添加失败, 删除锁
		return 0, status.Errorf(codes.InvalidArgument, "添加元素失败")
	}
	return result, nil
}

func (s *RedisServer) GenerateGroupSpecialCodes(digit int64, data []interface{}) (int64, error) {
	// 验证位数
	if _, err := s.SetGroupSpecialCodeDigit(digit); err != nil {
		return 0, err
	}

	// 设置锁
	//lockKey := RedisKey{}.GetGeneratorUserCodeLockKey(digit)
	//if res := s.SetNX(lockKey, strconv.Itoa(int(digit)), 0); res == false {
	//	return 0, status.Errorf(codes.AlreadyExists, "当前位数已存在，key:%s", lockKey)
	//}

	// 生成集合
	key := RedisKey{}.GetGroupSpecialCodesKey()
	result, _ := global.Redis.SAdd(context.Background(), key, data).Result()
	if result <= 0 {
		//_ = s.Delete(lockKey, strconv.Itoa(int(digit))) // 添加失败, 删除锁
		//return 0, status.Errorf(codes.InvalidArgument, "添加特殊元素失败")
	}
	return result, nil
}

func (s *RedisServer) RandomGroupCodes(count int64) ([]string, error) {
	key := RedisKey{}.GetGroupCodesKey()
	realCount := s.GetGroupCodesCount()
	if count > realCount {
		count = realCount
	}
	return global.Redis.SPopN(context.Background(), key, count).Result()
}

func (s *RedisServer) RandomGroupSpecialCodes(count int64) ([]string, error) {
	key := RedisKey{}.GetGroupSpecialCodesKey()
	realCount := s.GetGroupSpecialCodesCount()
	if count > realCount {
		count = realCount
	}
	return global.Redis.SPopN(context.Background(), key, count).Result()
}

func (s *RedisServer) GetGroupCodesCount() int64 {
	key := RedisKey{}.GetGroupCodesKey()
	res, _ := global.Redis.SCard(context.Background(), key).Result()
	return res
}

func (s *RedisServer) GetGroupSpecialCodesCount() int64 {
	key := RedisKey{}.GetGroupSpecialCodesKey()
	res, _ := global.Redis.SCard(context.Background(), key).Result()
	return res
}
