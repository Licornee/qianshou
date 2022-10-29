package dao

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "time"
)

var Db *gorm.DB

func Init() {
    var err error
    dsn := "host=localhost user=yindongpeng dbname=qianshou port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    // 获取通用数据库对象 sql.DB ，然后使用其提供的功能
    sqlDB, err := Db.DB()

    // SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
    sqlDB.SetMaxIdleConns(10)

    // SetMaxOpenConns 设置打开数据库连接的最大数量。
    sqlDB.SetMaxOpenConns(100)

    // SetConnMaxLifetime 设置了连接可复用的最大时间。
    sqlDB.SetConnMaxLifetime(time.Hour)

    if err != nil {
        fmt.Printf("PostgreSQL connect error %v\n", err)
    }

    if Db.Error != nil {
        fmt.Printf("Database error %v", Db.Error)
    }
}
