package conf

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "root:000719@tcp(127.0.0.1:3306)/minio_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// var db *gorm.DB // 全局变量，用于存储数据库连接
// func InitDB() error {
// 	dsn := "root:000719@tcp(127.0.0.1:3306)/minio_demo?charset=utf8mb4&parseTime=True&loc=Local"

// 	var err error
// 	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	// 设置连接池配置
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		return err
// 	}
// 	sqlDB.SetMaxIdleConns(10)  // 设置连接池中的最大空闲连接数
// 	sqlDB.SetMaxOpenConns(100) // 设置连接池中的最大打开连接数

// 	return nil
// }
