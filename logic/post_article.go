package logic

import (
	"go.uber.org/zap"
	"goweb/dao/mysql"
	"goweb/dao/redis"
	"goweb/models"
	"goweb/pkg/snowflake"
)

func PostArticle(article *models.PostArticle) (err error) {
	//利用雪花算法生成PostID
	article.PostID = snowflake.GenID()
	//存进数据库
	err = mysql.PostArticle(article)
	if err != nil {
		return err
	}
	err = redis.CreatePost(article.PostID)
	return
}

func GetArticle(id int64) (*models.PostArticleSDetail, error) {
	post, err := mysql.GetArticle(id)
	if err != nil {
		zap.L().Error("GetArticle(id int64) failed")
		return nil, err
	}
	//fmt.Println(post)

	//根据作者id查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("GetArticle(id int64)里面的mysql.GetUserByID(post.AuthorID) failed")
		return nil, err
	}

	//根据communityID查询社区详情
	community, err := mysql.GetCommunityDetailList(post.CommunityID)
	if err != nil {
		zap.L().Error("GetArticle(id int64)里面的mysql.GetCommunityDetial(post.CommunityID) failed")
		return nil, err
	}

	data := models.PostArticleSDetail{
		AuthorName:   user.Username,
		PostArticle:  post,
		GetCommunity: community,
	}
	//return mysql.GetArticle(id)
	return &data, nil
}

func GetArticleList(pagesize int64, offset int64) ([]*models.PostArticleSDetail, error) {
	posts, err := mysql.GetArticleList(pagesize, offset)
	if err != nil {
		return nil, err
	}
	data := make([]*models.PostArticleSDetail, 0, len(posts))
	for _, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("logic.GetArticle.range posts user failed")
			return nil, err
		}

		//根据communityID查询社区详情
		community, err := mysql.GetCommunityDetailList(post.CommunityID)
		if err != nil {
			zap.L().Error("logic.GetArticle.range posts community failed")
			return nil, err
		}

		articledetail := &models.PostArticleSDetail{
			AuthorName:   user.Username,
			PostArticle:  post,
			GetCommunity: community,
		}
		data = append(data, articledetail)
	}
	return data, nil
}

//
func GetArticleLists(p *models.ParamArticleLists) ([]*models.PostArticleSDetail, error) {
	//2.去redis查询id列表
	ids, err := redis.GetPostIdsOrder(p)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return nil, nil
	} //3.根据id去数据库中查询帖子详情
	posts := mysql.GetPostListByIds(ids)
	data := make([]*models.PostArticleSDetail, 0, len(posts))

	//提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	for idx, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("logic.GetArticle.range posts user failed")
			return nil, err
		}

		//根据communityID查询社区详情
		community, err := mysql.GetCommunityDetailList(post.CommunityID)
		if err != nil {
			zap.L().Error("logic.GetArticle.range posts community failed")
			return nil, err
		}

		articledetail := &models.PostArticleSDetail{
			AuthorName:   user.Username,
			VoteNum:      voteData[idx],
			PostArticle:  post,
			GetCommunity: community,
		}
		data = append(data, articledetail)
	}
	return nil, nil
}
