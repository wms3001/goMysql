package goMysql

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGoMysql_Connect(t *testing.T) {
	goMysql := GoMysql{}
	goMysql.Addr = "192.168.4.81"
	goMysql.Port = "3306"
	goMysql.User = "root"
	goMysql.Pass = "MyNewPass4!@#"
	goMysql.Db = "test"
	goMysql.MaxOpen = 20
	goMysql.MaxIdle = 20
	goMysql.MaxLifetime = 10
	resp := goMysql.Connect()
	defer goMysql.Close()
	jsonr, _ := json.Marshal(resp)
	fmt.Println(string(jsonr))
}

func TestGoMysql_Exec(t *testing.T) {
	goMysql := GoMysql{}
	goMysql.Addr = "192.168.4.81"
	goMysql.Port = "3306"
	goMysql.User = "root"
	goMysql.Pass = "MyNewPass4!@#"
	goMysql.Db = "test"
	goMysql.MaxOpen = 20
	goMysql.MaxIdle = 20
	goMysql.MaxLifetime = 10
	goMysql.Sql = "create table tttttt(    id   int auto_increment,    name varchar(20) null,    constraint tttttt_pk       primary key (id))"
	goMysql.Connect()
	resp := goMysql.Exec()
	defer goMysql.Close()
	jsonr, _ := json.Marshal(resp)
	fmt.Println(string(jsonr))
}

type Tttttt struct {
	Id    int64   `json:"id"`
	Test  string  `json:"test"`
	Total int64   `json:"total"`
	Price float64 `json:"price""`
}

func TestGoMysql_Select(t *testing.T) {
	goMysql := GoMysql{}
	goMysql.Addr = "192.168.4.81"
	goMysql.Port = "3306"
	goMysql.User = "root"
	goMysql.Pass = "MyNewPass4!@#"
	goMysql.Db = "test"
	goMysql.MaxOpen = 20
	goMysql.MaxIdle = 20
	goMysql.MaxLifetime = 10
	goMysql.Sql = "select * from wms"
	goMysql.Connect()
	defer goMysql.Close()
	resp := goMysql.Select()
	//jsonr, _ := json.Marshal(resp)
	//fmt.Println(string(jsonr))
	var ttt []Tttttt
	fmt.Println(resp.Data)
	json.Unmarshal([]byte(resp.Data), &ttt)
	fmt.Println(ttt)

}
