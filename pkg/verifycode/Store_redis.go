package verifycode

import (
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 verifycode.Store interface 的 Set 方法
// 保存验证码
func (s *RedisStore) Set(id string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))
	// 本地环境方便调试
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPrefix+id, value, ExpireTime)

}

// 获取验证码
func (s *RedisStore) Get(id string, clear bool) string {
	id = s.KeyPrefix + id
	val := s.RedisClient.Get(id)
	if clear {
		s.RedisClient.Del(id)
	}
	return val
}

// 检查验证码
func (s *RedisStore) Verify(id string, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
