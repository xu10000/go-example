package main

import (
	"database/sql"
	"fmt"

	// init 初始化了database/sql库中的全局变量，所以下面可直接调用！
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ApiDB?charset=utf8") //第一个参数数驱动名
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("------ db ", db)
}
func main() {

}
