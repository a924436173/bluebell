package redis

// redis key

// redis key 尽量使用命名空间的方式,方便查询和拆分

const (
	KeyPrefix          = "bluebell:"   // 项目key前缀
	KeyPostTimeZSet    = "post:time"   // zset;帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  // zset;帖子及投票的分数
	KeyPostVotedZSetPF = "post:voted:" // zset;记录用户及投票的类型;参数是 post id

	KeyCommunitySetPF = "community:" // set;保存每个分区下帖子的id

)

// getRedisKey 给 rediskey 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
