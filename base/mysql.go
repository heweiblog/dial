package base

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// mysql 拨测 addr:ip:port 连接超时2s 查询超时2s 返回时延
func Mysql(addr, user, pass, dbname, method string) int64 {
	t := time.Now()
	par := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&timeout=2s", user, pass, addr, dbname)

	db, err := sql.Open("mysql", par)
	if err != nil {
		return 0
	}
	defer db.Close()

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	_, err = db.QueryContext(ctx, method)
	if err != nil {
		return 0
	}

	return time.Since(t).Nanoseconds() / 1000
}
