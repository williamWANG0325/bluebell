package redis

import (
	"errors"
	"fmt"
	"math"
	"time"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

/* 投票的几种情况：
direction=1时，有两种情况：
	1.之前没有投过票，现在投赞成票
	2.之前投反对票，现在改投赞成票
direction=0时，有两种情况：
	1.之前投过赞成票，现在要取消投票
	2. 之前投过反对票，现在要取消投票
direction=-1时，有两种情况：
	1. 之前没有投过票，现在投反对票
	2. 之前投赞成票，现在改投反对票

投票的限制：
每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
	1．到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekInSecond = 7 * 24 * 3600
	scorePreVote    = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postId int64) error {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	rdb.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	// 帖子分数
	rdb.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) (err error) {
	// 1. 判断投票限制
	// 获取发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	zap.L().Error(fmt.Sprintf("\n%v\n%v\n", postTime, postID))
	if float64(time.Now().Unix())-postTime > oneWeekInSecond {
		return ErrVoteTimeExpire
	}
	// 2. 更新贴子的分数
	// 先查当前用户给当前帖子的投票记录
	op := rdb.ZScore(getRedisKey(KeyPostVotedZSerPF+postID), userID).Val()
	if value == op {
		return ErrVoteRepeated
	}

	diff := math.Abs(op - value)
	var dir float64
	if value > op {
		dir = 1
	} else {
		dir = -1
	}
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePreVote, postID)

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSerPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSerPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipeline.Exec()
	return
}
