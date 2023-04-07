package models

import "time"

type PostArticle struct {
	Status      int8      `json:"status" gorm:"status"`
	PostID      int64     `json:"post_id" gorm:"post_id"`
	ID          int64     `json:"id" gorm:"id"`
	AuthorID    int64     `json:"author_id" gorm:"author_id"`
	CommunityID int64     `json:"community_id" gorm:"community_id"`
	Title       string    `json:"title" gorm:"title"`
	Content     string    `json:"content" gorm:"content"`
	CreateTime  time.Time `json:"create_time" gorm:"create_time"`
}

// PostArticleSDetail 获取帖子详情信息结构体
type PostArticleSDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum    int64  `json:"vote_num"`
	*PostArticle
	*GetCommunity
}
