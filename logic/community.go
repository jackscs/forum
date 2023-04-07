package logic

import (
	"goweb/dao/mysql"
	"goweb/models"
)

func GetCommunityList() ([]models.Community, error) {
	//查询数据库找到所有的community并返回
	data, err := mysql.GetCommunityList()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetCommunityDetailList(id int64) (*models.GetCommunity, error) {
	return mysql.GetCommunityDetailList(id)
}
