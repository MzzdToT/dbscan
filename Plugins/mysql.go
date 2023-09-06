package Plugins

import (
    "database/sql"
    "fmt"
    "io/ioutil"
    "strings"
    _ "github.com/go-sql-driver/mysql"
)

// MysqlScan 尝试使用提供的用户名和密码爆破 MySQL 数据库
func MysqlScan(ip string, userFile string, passFile string, dbPort int) error {
    // 读取账号文件中的内容
    userContent, err := ioutil.ReadFile(userFile)
    if err != nil {
        return err
    }
    users := strings.Split(string(userContent), "\n")

    // 读取密码文件中的内容
    passContent, err := ioutil.ReadFile(passFile)
    if err != nil {
        return err
    }
    passes := strings.Split(string(passContent), "\n")

    // 构建 MySQL 连接字符串
    dataSourceName := fmt.Sprintf("root:password@tcp(%s:%d)/mysql", ip, dbPort)

    // 使用账号和密码组合进行破解
    for _, user := range users {
        for _, pass := range passes {
            if err := MysqlConn(dataSourceName, user, pass); err == nil {
                fmt.Printf("爆破成功! ：%s:%s\n", user, pass)
                return nil
            } else {
                fmt.Println(fmt.Sprintf("正在测试：%s:%s", user, pass))
            }
        }
    }

    return fmt.Errorf("未成功爆破出账号密码")
}

// MysqlConn 尝试建立 MySQL 数据库连接
func MysqlConn(dataSourceName string, user string, pass string) error {
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
