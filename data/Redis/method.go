package RedisDB

import (
	"context"
	"time"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

func Set(c *redis.Client, ctx context.Context, key string, value interface{}, expiration time.Duration) error {
    p, err := json.Marshal(value)
    if err != nil {
       return err
    }
    return c.Set(ctx, key, p, expiration).Err()
}

func Get(c *redis.Client, ctx context.Context, key string, value interface{}) error {
	b, err := c.Get(ctx, key).Result()
   if err != nil {
      return err
   }
    return json.Unmarshal([]byte(b), value)
}