package redis

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Set(key string, value any, expiration time.Duration) error {
	if redisInstance == nil {
		return fmt.Errorf("redis is not initialized")
	}

	err := redisInstance.Set(ctx, key, value, expiration).Err()
	if err != nil {
		fmt.Printf("Error setting Redis key %s: %v\n", key, err)
		return err
	}
	return nil
}

func Get(key string) (string, error) {
	if redisInstance == nil {
		return "", fmt.Errorf("redis is not initialized")
	}

	val, err := redisInstance.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key %s does not exist", key)
	} else if err != nil {
		fmt.Printf("Error getting Redis key %s: %v\n", key, err)
		return "", err
	}
	return val, nil
}

func Delete(key string) error {
	if redisInstance == nil {
		return fmt.Errorf("redis is not initialized")
	}

	err := redisInstance.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Error deleting Redis key %s: %v\n", key, err)
		return err
	}
	return nil
}

func DeletePattern(pattern string) (int64, error) {
	if redisInstance == nil {
		return 0, fmt.Errorf("redis is not initialized")
	}

	keys, err := redisInstance.Keys(ctx, pattern).Result()
	if err != nil {
		fmt.Printf("Error finding keys matching pattern %s: %v\n", pattern, err)
		return 0, err
	}

	if len(keys) == 0 {
		return 0, nil
	}

	deleted, err := redisInstance.Del(ctx, keys...).Result()
	if err != nil {
		fmt.Printf("Error deleting Redis keys: %v\n", err)
		return 0, err
	}
	return deleted, nil
}

func Append(key string, value string) error {
	if redisInstance == nil {
		return fmt.Errorf("redis is not initialized")
	}

	err := redisInstance.Append(ctx, key, value).Err()
	if err != nil {
		fmt.Printf("Error appending to Redis key %s: %v\n", key, err)
		return err
	}
	return nil
}

func Expire(key string, expiration time.Duration) error {
	if redisInstance == nil {
		return fmt.Errorf("redis is not initialized")
	}

	err := redisInstance.Expire(ctx, key, expiration).Err()
	if err != nil {
		fmt.Printf("Error setting expiration for Redis key %s: %v\n", key, err)
		return err
	}
	return nil
}
