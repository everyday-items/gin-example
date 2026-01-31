package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/everyday-items/gin-example/library/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Setup 初始化数据库连接
func Setup() {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			setting.MysqlSetting.Username,
			setting.MysqlSetting.Password,
			setting.MysqlSetting.Host,
			setting.MysqlSetting.Port,
			setting.MysqlSetting.Database,
			setting.MysqlSetting.Charset,
		)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to get database instance: %v", err)
		}

		// 设置连接池参数
		sqlDB.SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)
		sqlDB.SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)

		log.Println("Database connection established successfully")
	})
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized. Call Setup() first.")
	}
	return db
}

// Close 关闭数据库连接
func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
