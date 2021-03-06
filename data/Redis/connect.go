package DatabaseRedisDB

import (
	"context"
	"github.com/go-redis/redis/v8"
)


func Init(RedisURL, RedisPassword string) ( *redis.Client, error ) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     RedisURL,
		Password: RedisPassword,
		DB:       0,  
	})		

	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		_ = RedisClient.Close()
		return nil, err
	}

	return RedisClient, nil
}

func Disconnect( RedisClient *redis.Client ) error {
	if err := RedisClient.Close(); err != nil {
		return err
	} else {
		return nil
	}

}
