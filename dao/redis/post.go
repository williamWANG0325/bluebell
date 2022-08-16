package redis

import (
	"bluebell/models"

	"github.com/go-redis/redis"
)

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}

	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostVoteData(ids []string) (data []int64, err error) {

	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSerPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
