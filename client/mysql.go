package client

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"model-service/config"
)

var db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.User,
		config.Config.Mysql.Pass,
		config.Config.Mysql.Addr,
		config.Config.Mysql.DBName,
	)
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}
	configConnectionPool()
}
func configConnectionPool() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
func GetDB() *gorm.DB {
	return db
}
