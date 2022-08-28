package redis

import (
	"go.uber.org/zap"
	"strconv"
)

func AddToRedisSet(userid, fans int64) {
	_, err := client.SAdd(ctx, KeyFollowing+strconv.Itoa(int(userid)), fans).Result()
	if err != nil {
		zap.L().Error("关注集合缓存到redis失败")
		return
	}
}
