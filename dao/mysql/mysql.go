package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"goweb/settings"
)

// gorm 2.0连接版本
func NewDBEngine(databaseSetting *settings.MysqlConfig) (*gorm.DB, error) {

	s := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
	)
	//fmt.Println(dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		zap.L().Error(" mysql中NewDBEngine方法初始化数据库失败")
		return nil, err
	}
	return db, nil
}
