package redis

//redis key

//redis key 注意使用命名规范

const (
	Prefix             = "bluebell"
	KeyPostTimeZSet    = "post.time"  //zset:帖子及发帖时间
	KeyPostScoreZSet   = "post.score" //zset:帖子及投票分数
	KeyPostVotedZSetPF = "post.voted" //zset:记录用户投票类型；参数是post id
)

func getRedisKey(key string) string {
	return Prefix + key
}
