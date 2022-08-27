package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := client.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet),
		redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: postID,
		})

	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet),
		redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: postID,
		})
	_, err := pipeline.Exec()

	return err

}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票限制
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	//2. 更新帖子的分数
	ov := client.ZScore(getRedisKey(KeyPostScoreZSet+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePreVote, postID)

	//3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postID), postID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}

	_, err := pipeline.Exec()
	return err
}
