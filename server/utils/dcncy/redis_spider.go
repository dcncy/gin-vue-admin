package dcncy

import (
	"context"
	"time"

	"github.com/dcncy/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

var SPIDER_REDIS *SpiderRedisStore

func init() {
	SPIDER_REDIS = &SpiderRedisStore{
		Expiration: time.Hour * 24 * 30,
		PreKey:     "SPIDER_",
		Context:    context.TODO(),
	}
}

type SpiderRedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *SpiderRedisStore) Set(id string, value string) error {
	err := global.GVA_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *SpiderRedisStore) Get(key string, clear bool) (string, bool) {
	val, err := global.GVA_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))
		return "", false
	}
	if clear {
		err := global.GVA_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.GVA_LOG.Error("RedisStoreClearError!", zap.Error(err))
			return "", false
		}
	}
	return val, true
}

func (rs *SpiderRedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v, _ := rs.Get(key, clear)
	return v == answer
}
