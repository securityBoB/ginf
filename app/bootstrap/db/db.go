package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

func InitDB() (*sql.DB, error) {
	viper.SetConfigFile("config/app.toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}

	// 从配置文件中获取数据库信息
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbName := viper.GetString("database.dbname")
	dbuser := viper.GetString("database.username")
	dbPass := viper.GetString("database.password")
	dbCharset := viper.GetString("database.charset")

	// 拼接数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbuser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		dbCharset,
	)

	// 初始化数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db, nil
}
