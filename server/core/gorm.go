package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/global"
	"time"
)

func Gorm() *gorm.DB {
	return InitGorm()
}

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warn("mysql config is empty")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}
	global.MysqlLog = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.Log.Error("mysql [%s] connect error", dsn)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxConnections)
	sqlDB.SetConnMaxIdleTime(time.Hour * 4)
	return db
}
