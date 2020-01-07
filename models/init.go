package models

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	var err error
	//加载配置
	cfg, err := goconfig.LoadConfigFile("./config/env.ini")
	if err != nil {
		panic("read config error")
	}

	db_driver, err := cfg.GetValue("db-dev", "db_driver")
	if err != nil {
		panic("read config error")
	}
	db_dsn, err := cfg.GetValue("db-dev", "db_dsn")
	if err != nil {
		panic("read config error")
	}

	DB, err = gorm.Open(db_driver, db_dsn)

	if err != nil {
		fmt.Printf("database connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	migration()

}
