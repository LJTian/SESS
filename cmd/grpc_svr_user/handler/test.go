package handler

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"SESS/cmd/grpc_svr_user/config"
)

var (
	dbcfg = config.DBInfo{
		IP:       "127.0.0.1",
		Port:     23306,
		Name:     "user_svr",
		UserName: "root",
		PassWord: "123456",
	}

	DB *gorm.DB
)

func InitDB() {
	c := dbcfg
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.UserName, c.PassWord, c.IP, c.Port, c.Name)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // 禁用彩色打印
		},
	)

	// 全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}
