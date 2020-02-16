package data

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"

	_ "github.com/lib/pq"
)

const (
	//host     = "pgm-bp1bhd284mgr043efo.pg.rds.aliyuncs.com" // 外网
	host = "pgm-bp1bhd284mgr043e129210.pg.rds.aliyuncs.com"
	port     = 1433
	user     = "devdb"
	password = "db@123456"
	dbName   = "wuhan"
)

var db *xorm.Engine

func InitDB() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbName)
	fmt.Println(psqlInfo)
	//格式
	db, err = xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	//db.ShowSQL(true)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")
}
