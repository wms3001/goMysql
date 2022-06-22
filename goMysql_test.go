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
	resp := goMysql.Connect()
	jsonr, _ := json.Marshal(resp)
	fmt.Println(string(jsonr))
}
