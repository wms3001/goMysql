package goMysql

import (
	"database/sql"
	"encoding/json"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/goTool"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type GoMysql struct {
	Addr        string
	Port        string
	User        string
	Pass        string
	Db          string
	Sql         string
	MaxOpen     int
	MaxIdle     int
	MaxLifetime time.Duration
	Conn        *sql.DB
	Stmt        *sql.Stmt
}

func (goMysql *GoMysql) Connect() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	dsn := goMysql.User + ":" + goMysql.Pass + "@tcp(" + goMysql.Addr + ":" + goMysql.Port + ")/" + goMysql.Db + "?charset=utf8mb4&parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		db.SetConnMaxLifetime(time.Minute * goMysql.MaxLifetime)
		db.SetMaxOpenConns(goMysql.MaxOpen)
		db.SetMaxIdleConns(goMysql.MaxIdle)
		resp.Code = 1
		resp.Message = "connected"
		goMysql.Conn = db
	}
	return resp
}

func (goMysql *GoMysql) Close() {
	goMysql.Conn.Close()
}

func (goMysql *GoMysql) Exec() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	re, err := goMysql.Conn.Exec(goMysql.Sql)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		lastId, _ := re.LastInsertId()
		roweffct, _ := re.RowsAffected()
		data := make(map[string]int64)
		data["lastInsertId"] = lastId
		data["rowsAffected"] = roweffct
		dataType, _ := json.Marshal(data)
		dataString := string(dataType)
		resp.Data = dataString
	}
	return resp
}

func (goMysql *GoMysql) Prepare() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	stmt, err := goMysql.Conn.Prepare(goMysql.Sql)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		goMysql.Stmt = stmt
	}
	return resp
}

func (goMysql *GoMysql) Select() *goCommon.Resp {
	var goTool = goTool.GoTool{}
	var resp = &goCommon.Resp{}
	rows, err := goMysql.Conn.Query(goMysql.Sql)
	defer rows.Close()
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		columns, _ := rows.Columns()
		count := len(columns)
		values := make([]interface{}, count)
		scanArgs := make([]interface{}, count)
		for i := range values {
			scanArgs[i] = &values[i]
		}
		var ttt []map[string]interface{}
		for rows.Next() {
			err := rows.Scan(scanArgs...)
			if err != nil {
				resp.Code = -1
				resp.Message = err.Error()
				break
			}
			mm := make(map[string]interface{})
			for i, v := range values {
				nn := goTool.Strval(v)
				mm[columns[i]] = nn
				//x := v.([]byte)
				//if nx, ok := strconv.ParseFloat(string(x), 64); ok == nil {
				//	//masterData[columns[i]] = append(masterData[columns[i]], nx)
				//	mm[columns[i]] = nx
				//} else if b, ok := strconv.ParseBool(string(x)); ok == nil {
				//	//masterData[columns[i]] = append(masterData[columns[i]], b)
				//	mm[columns[i]] = b
				//} else if "string" == fmt.Sprintf("%T", string(x)) {
				//	//masterData[columns[i]] = append(masterData[columns[i]], string(x))
				//	mm[columns[i]] = string(x)
				//	//} else if reflect.TypeOf(x) {
				//} else {
				//	fmt.Printf("Failed on if for type %T of %v\n", x, x)
				//}
			}
			ttt = append(ttt, mm)
		}
		dataType, _ := json.Marshal(ttt)
		dataString := string(dataType)
		resp.Code = 1
		resp.Message = "success"
		resp.Data = dataString
	}
	return resp
}
