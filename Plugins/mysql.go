package Plugins

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
    
)


// MysqlScan 尝试使用提供的用户名和密码爆破 MySQL 数据库
func MysqlScan(ip string, user string, pass string, dbPort int) error {

    // 构建 MySQL 连接字符串
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%v:%v)/mysql?charset=utf8", user, pass, ip, dbPort)
    fmt.Println(user,pass,dataSourceName)
    if err := MysqlConn(dataSourceName); err == nil {
        fmt.Printf("爆破成功! ：%v:%v %v:%v\n",ip ,dbPort , user, pass)
        return nil
    } else {
        return err
    }
}


// MysqlConn 尝试建立 MySQL 数据库连接
func MysqlConn(dataSourceName string) error {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return err
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        return err
    }

    return nil
}

