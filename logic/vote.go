package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"
)

// 投票功能

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}
