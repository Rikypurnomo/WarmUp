package cache

import (
	"context"
	"time"

	"github.com/Rikypurnomo/warmup/config"
	"github.com/Rikypurnomo/warmup/pkg/logger"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	ClientRedis *redis.Client
	ContextTim  = context.Background()
	tracer      = otel.Tracer("github.com/Rikypurnomo/warmup/internal/api/handlers/")
)

func InitCache() *redis.Client {
	ClientRedis = redis.NewClient(&redis.Options{
		Addr:     config.RedisToAddr(),
		Password: config.PasswordRedis(),
		DB:       0,
	})

	_, err := ClientRedis.Ping(context.Background()).Result()
	if err != nil {
		panic("error" + err.Error())
	}
	logger.Debug("Connect to redis")
	return ClientRedis
}

func CloseConnectionCache() {
	logger.Debug("Close redis connection")
	ClientRedis.Close()
}

func SetKeyArray(context context.Context, key string, i []byte, ttl time.Duration) error {
	context, span := tracer.Start(context, "Redis.SetKeyArray")
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key("key"),
		Value: attribute.StringValue(key),
	})
	if ttl == 0 {
		ClientRedis.Set(context, key, i, 0)
	} else {
		ClientRedis.Set(context, key, i, ttl)
	}

	return ClientRedis.Get(context, key).Err()
}

func SetKey(context context.Context, key string, i interface{}, ttl time.Duration) *redis.StringCmd {
	context, span := tracer.Start(context, "Redis.SetKey")
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key("key"),
		Value: attribute.StringValue(key),
	})
	if ttl == 0 {
		ClientRedis.Set(context, key, i, 0)
	} else {
		ClientRedis.Set(context, key, i, ttl)
	}

	return ClientRedis.Get(context, key)
}

func GetKey(context context.Context, key string) *redis.StringCmd {
	context, span := tracer.Start(context, "Redis.GetKey")
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key("key"),
		Value: attribute.StringValue(key),
	})
	return ClientRedis.Get(context, key)
}

func FindKey(context context.Context, key string) *redis.StringSliceCmd {
	context, span := tracer.Start(context, "Redis.FindKey")
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key("key"),
		Value: attribute.StringValue(key),
	})
	return ClientRedis.Keys(context, key)
}

func Scan(match string) (keys []string) {
	idb := ClientRedis.Scan(ContextTim, 0, match, 0).Iterator()
	for idb.Next(ContextTim) {
		keys = append(keys, idb.Val())
	}
	return keys
}

func RemoveKey(context context.Context, key string) error {
	context, span := tracer.Start(context, "Redis.RemoveKey")
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   attribute.Key("key"),
		Value: attribute.StringValue(key),
	})
	err := ClientRedis.Del(context, key).Err()
	return err
}

func RemoveKeys(context context.Context, key ...string) error {
	err := ClientRedis.Del(context, key...).Err()
	return err
}

func HasKey(context context.Context, key string) bool {
	return ClientRedis.Exists(context, key).Val() == 1
}
