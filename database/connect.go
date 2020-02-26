package database

import (
	"app/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/xormplus/core"
	"time"
	//"github.com/xormplus/xorm"
	"xorm.io/xorm"
)

var (
	db    *xorm.Engine
	param = "parseTime=true"
)

func init() {
	// connect database sql.Open("mysql", "user:password@/dbname")
	var err error
	db, err = xorm.NewEngine(setting.DBs["main"].Driver,
		fmt.Sprintf("%s:%s@/%s?%s", setting.DBs["main"].User, setting.DBs["main"].Password, setting.DBs["main"].Name, setting.DBs["main"].Param))
	if err != nil {
		panic(err)
	}

	// optimize option
	db.SetMaxOpenConns(setting.DBs["main"].Option["SetMaxOpenConns"])
	db.SetMaxIdleConns(setting.DBs["main"].Option["SetMaxIdleConns"])
	db.SetConnMaxLifetime(time.Duration(setting.DBs["main"].Option["SetConnMaxLifetime"]) * time.Second)

	if setting.Servers["main"].RunMode == "debug" {
		db.ShowSQL(true)
		//db.Logger().SetLevel(core.LOG_DEBUG)
	}
	fmt.Print("init")
}

func Close() error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
