package queue

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(url string) *RedisQueue {
	opt, _ := redis.ParseURL(url)
	client := redis.NewClient(opt)

	return &RedisQueue{client: client}
}

func (q *RedisQueue) Push(queue string, data string) error {
	return q.client.LPush(ctx, queue, data).Err()
}

func (q *RedisQueue) Pop(queue string) (string, error) {
	return q.client.RPop(ctx, queue).Result()
}
