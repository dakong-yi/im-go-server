package db

import (
	"fmt"

	"github.com/dakong-yi/im-go-server/config"
	"github.com/dakong-yi/im-go-server/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	mysqlConfig := config.Cfg.MysqlCfg
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database, mysqlConfig.Charset,
		mysqlConfig.ParseTime, mysqlConfig.TimeZone)
	DB, err = gorm.Open(mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(fmt.Sprintf("创建mysql客户端失败: %v,%s", DB, err))
	}
	// Perform auto-migration
	DB.AutoMigrate(&model.Conversation{}, &model.Message{}, &model.UserConversation{})

	// Return nil if there is no error
	return nil
}
