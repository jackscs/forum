package models

import "time"

//获取到列表
type Community struct {
	ID   int64  `json:"id" gorm:"column:community_id"`
	Name string `json:"name" gorm:"column:community_name"`
}

//通过id获取到详细列表
type GetCommunity struct {
	Id           int64     `json:"id" gorm:"column:community_id`
	Community    string    `json:"community" gorm:"column:community_name"`
	Introduction string    `json:"introduction" gorm:"introduction"`
	CreateTime   time.Time `json:"createtime" gorm:"create"`
}
