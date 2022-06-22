package goMysql

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type GoMysql struct {
	Addr string
	Port string
	User string
	Pass string
	Db   string
	Conn *sql.DB
}

func (goMysql *GoMysql) Connect() *Resp {
	var resp = &Resp{}
	dsn := goMysql.User + ":" + goMysql.Pass + "@/tcp(" + goMysql.Addr + ":" + goMysql.Port + ")/" + goMysql.Db + "?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "connected"
		goMysql.Conn = db
	}
	return resp
}
