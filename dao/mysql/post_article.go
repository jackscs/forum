package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"goweb/global"
	"goweb/models"
)

func PostArticle(article *models.PostArticle) (err error) {
	//将文章等信息存进数据库
	err = global.DBEngine.Debug().Table("article").Create(article).Error
	fmt.Println(article)
	if err != nil {
		zap.L().Error(" Postarticle()文章入库失败")
		return err
	}
	return
}

func GetArticle(id int64) (*models.PostArticle, error) {
	article := new(models.PostArticle)
	err := global.DBEngine.Table("article").Where("id=?", id).Find(&article).Error
	if err != nil {
		return &models.PostArticle{}, err
	}
	return article, nil
}

// GetArticleList 查询帖子函数
func GetArticleList(pagesize int64, offset int64) ([]*models.PostArticle, error) {
	pagesizes := int(pagesize)
	offsets := int(offset)
	articles := make([]*models.PostArticle, 0, 2)
	err := global.DBEngine.Table("article").Limit(pagesizes).Offset(offsets).Find(&articles).Error
	if err != nil {
		zap.L().Error("mysql.GetArticleList failed", zap.Error(err))
		return nil, err
	}
	return articles, nil
}

//根据给定id列表查询帖子数据
func GetPostListByIds(ids []string) []*models.PostArticle {
	var article []*models.PostArticle
	global.DBEngine.Table("post").Where("id in", ids).Order("post_id desc").Find(&article)
	return article
}
