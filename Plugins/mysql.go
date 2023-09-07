package Plugins

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "strings"
    _ "github.com/go-sql-driver/mysql"
    
)

// MysqlScan 尝试使用提供的用户名和密码爆破 MySQL 数据库
func MysqlScan(ip string, user string, pass string, dbPort int) error {
    fmt.Println("ok")
    // 构建 MySQL 连接字符串
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql", user, pass,ip, dbPort)
    fmt.Println(dataSourceName)
    if err := MysqlConn(dataSourceName); err == nil {
        fmt.Printf("爆破成功! ：%s:%d %s:%s\n",ip ,dbPort , user, pass)
        return nil
    } else {
        fmt.Printf("正在测试：%s:%d %s:%s", ip,dbPort , user, pass)
        return fmt.Errorf("连接失败: %v", err)
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
