package redis

import (
	"go.uber.org/zap"
	"strconv"
)

// FollowSet 缓存到关注集合
func FollowSet(key string, userid, fans int64) error {
	_, err := client.SAdd(ctx, key+strconv.Itoa(int(userid)), fans).Result()
	if err != nil {
		zap.L().Error("关注集合缓存到redis失败")
		return err
	}
	return err
}

// FollowRemove 移除关注集合
func FollowRemove(key string, userid, fans int64) error {
	_, err := client.SRem(ctx, key+strconv.Itoa(int(userid)), fans).Result()
	if err != nil {
		zap.L().Error("移除关注缓存redis失败")
		return err
	}
	return err
}

// ReadFollowSet 查询关注集合
func ReadFollowSet(userid int64) ([]string, error) {
	res, err := client.SMembers(ctx, KeyFollowing+strconv.Itoa(int(userid))).Result()
	if err != nil {
		zap.L().Error("重新")
		return nil, err
	}
	return res, err
}
