package mysql

import (
	"errors"
	"go.uber.org/zap"
	"goweb/global"
	"goweb/models"
)

func GetCommunityList() (data []models.Community, err error) {
	var typeid []models.Community
	err = global.DBEngine.Table("community").Find(&typeid).Error
	if err != nil {
		zap.L().Error("GetcommunityList failed")
		return nil, errors.New("获取参数列表失败")
	}
	return typeid, nil
}

func GetCommunityDetailList(id int64) (community *models.GetCommunity, err error) {
	var communityDetail models.GetCommunity
	err = global.DBEngine.Table("community").Where("community_id=?", id).Find(&communityDetail).Error
	if err != nil {
		zap.L().Error("GetCommunityDetailList failed")
		return nil, err
	}
	return &communityDetail, nil
}
