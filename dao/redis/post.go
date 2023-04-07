package redis

import (
	"github.com/go-redis/redis"
	"goweb/models"
)

func GetPostIdsOrder(p *models.ParamArticleLists) ([]string, error) {
	//从redis中获取id
	//1.根据用户请求中携带的order参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引的起点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//3.查询
	return client.ZRange(key, start, end).Result()
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//此写法不太好会影响性能
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	v1 := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v1)
	//}
	//return

	//使用pipeline一次发送多次请求较少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
