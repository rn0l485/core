package DatabaseRedisDB

import (
	"context"
	"time"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

func Set(c *redis.Client, ctx context.Context, key string, value interface{}, expiration ...time.Duration) error {
   p, err := json.Marshal(value)
   if err != nil {
     return err
   }

   var expire time.Duration
   if len(expiration) == 0 {
      expire = 0
   } else {
      expire = expiration[0]
   }

   return c.Set(ctx, key, p, expire).Err()
}

func Get(c *redis.Client, ctx context.Context, key string, value interface{}, raw ...bool) error {
   b, err := c.Get(ctx, key).Result()
   if err != nil {
      return err
   }

   if len(raw) == 0 || !raw[0] {
      return json.Unmarshal([]byte(b), value)
   }

   if v, ok := value.(*string); !ok {
      return json.Unmarshal([]byte(b), value)
   } else {
      *v = b
      return nil
   }
}

func Del(c *redis.Client, ctx context.Context, keys ...string) error {
   if result, err := c.Del(ctx, keys).Result(); err != nil {
      return err
   } else {
      return nil
   }
}