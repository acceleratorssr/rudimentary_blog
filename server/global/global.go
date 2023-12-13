package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"server/config"
)

var (
	// Config 此处需要类型记得是 *config.Config
	// Config应该是指针类型接收 /core/conf.go 中的结构体config
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
)
