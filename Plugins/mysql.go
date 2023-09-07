package Plugins

import (
    "errors"
    "fmt"
)

// MysqlScan 尝试使用提供的用户名和密码爆破 MySQL 数据库
func MysqlScan(ip string ,user string, pass string, dbPort int) error {
    // 在这里进行一些检查，如果发现错误情况，就返回一个错误
    if user == "" {
        return errors.New("用户名不能为空")
    }

    if pass == "" {
        return errors.New("密码不能为空")
    }

    if dbPort <= 0 {
        return errors.New("无效的端口号")
    }

    // 如果一切正常，返回 nil，表示没有错误
    fmt.Println("%s:%d user:%s pass:%s ",ip,dbPort,user,pass)
    return nil
}
