package logic

import (
	"goweb/dao/redis"
	"goweb/models"
	"strconv"
)

//投票功能
//简化版投票

//VotForPost 为帖子投票的函数
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	return redis.PostVote(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
