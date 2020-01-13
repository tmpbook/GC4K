package models

import (
	"fmt"

	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB

func init() {
	var err error
	//加载配置
	cfg, err := goconfig.LoadConfigFile("./config/env.ini")
	if err != nil {
		panic("read config error")
	}

	dbDriver, err := cfg.GetValue("DbDev", "DbDriver")
	if err != nil {
		panic("read config error")
	}
	dbDsn, err := cfg.GetValue("DbDev", "DbDsn")
	if err != nil {
		panic("read config error")
	}

	DB, err = gorm.Open(dbDriver, dbDsn)

	if err != nil {
		fmt.Printf("database connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}

	migration()

}
