package cache

import "strconv"

type RedisKey struct {
}

const prefix = "qvbilam:"

// GetGeneratorGroupCodeLockKey 用户 code 生成器
func (k RedisKey) GetGeneratorGroupCodeLockKey(count int64) string {
	return prefix + "generator:group:codes:lock:" + strconv.Itoa(int(count))
}

func (k RedisKey) GetGeneratorGroupCodeMaxDigit() string {
	return prefix + "generator:group:codes:ordinary:digit"
}

func (k RedisKey) GetGeneratorGroupSpecialCodeMaxDigit() string {
	return prefix + "generator:group:codes:special:digit"
}

// GetGroupCodesKey 获取用户code
func (k RedisKey) GetGroupCodesKey() string {
	return prefix + "group:codes:ordinary"
}

// GetGroupSpecialCodesKey 获取用户特殊code
func (k RedisKey) GetGroupSpecialCodesKey() string {
	return prefix + "group:codes:special"
}
