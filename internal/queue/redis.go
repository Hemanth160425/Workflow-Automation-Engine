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

    return &RedisQueue{
        client: client,
    }
}

func (q *RedisQueue) Push(queueName string, data string) error {
    return q.client.LPush(ctx, queueName, data).Err()
}


func (q *RedisQueue) Pop(queueName string) (string, error) {
    return q.client.RPop(ctx, queueName).Result()
}
