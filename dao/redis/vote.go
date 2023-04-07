package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

/* 投票的几种情况
direction=1时，有两种情况:
    1.之前没有投赞成票，现在投赞成票   -->更新分数和投票记录
    2.之前投反对票，现在投赞成票      -->更新分数和投票纪录
direction=0时，两种情况:
    1.之前投过赞成票,现在要取消投票    -->更新分数和投票纪录
    2.之前投过反对票，现在要取消投票   -->更新分数和投票纪录
direction=-1时,两种情况
    1.之前没有投过票，现在取消投票     -->更新分数和投票纪录
    2.之前投赞成票，现在改为反对票     -->更新分数和投票纪录

投票的限制:
每个帖子自发表之日一周之内允许用户投票，超过一个星期就不允许投票
    1.到期之后将redis中的赞成票和反对票保存进数据库中
    2.到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekSeconds = 7 * 24 * 3600
	scorePerVote   = 432 //每一票的分数
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepested   = errors.New("不允许重复投票")
)

func CreatePost(postID int64) error {

	//事务操作
	pipeline := client.TxPipeline()

	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	_, err := pipeline.Exec()
	return err
}

func PostVote(userID, postID string, value float64) error {
	//1.判断投票的限制
	//获取帖子的发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekSeconds {
		return ErrVoteTimeExpire
	}
	//2.更新帖子的分数
	//先查看当前用户给当前帖子的投票纪录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()

	//更新：防止重复投票
	if value == ov {
		return ErrVoteRepested
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)
	//3.记录用户为该帖子投票的
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
