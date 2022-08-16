package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"   // zset; 帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset; 帖子及分数
	KeyPostVotedZSerPF = "Post:voted:" // zset; 记录用户及投票类型； 参数是post_id
)

func getRedisKey(key string) (res string) {
	return KeyPrefix + key
}
