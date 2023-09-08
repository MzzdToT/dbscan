package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "github.com/spf13/viper" // 导入 YAML 解析库
    "github.com/MzzdToT/dbscan/Plugins"
    //"regexp"
)

func main() {
    configFile := flag.String("config", "", "Path to config file")
    singleIP := flag.String("t", "", "Specify a single IP")
    ipListFile := flag.String("T", "", "Path to IP list file")
    flag.Parse()

    // 如果未指定配置文件，则默认使用 config.yaml
    if *configFile == "" {
        *configFile = "config.yaml"
    }

    // 初始化配置解析器
    viper.SetConfigFile(*configFile)
    viper.SetConfigType("yaml")

    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("Error reading config file:", err)
        os.Exit(1)
    }

    mysqlType := viper.GetString("database.0.type")
    mysqlPort := viper.GetInt("database.0.port")

    /*oracleType := viper.GetString("database.1.type")
    oraclePort := viper.GetInt("database.1.port")*/

    //fmt.Println("MySQL Type:", mysqlType, "Port:", mysqlPort)
    /*fmt.Println("Oracle Type:", oracleType, "Port:", oraclePort)*/
    
    success := false
    //regex := regexp.MustCompile(`[\r\n]+$`)
    if *singleIP != "" {
        //fmt.Println(*singleIP)
        // 使用单个 IP 进行扫描
        switch mysqlType {
        case "mysql":
            users, err := readFile("./dic/mysql_user.txt")
            if err != nil {
                fmt.Println("Error reading user file:", err)
                os.Exit(1)
            }

            passwords, err := readFile("./dic/mysql_pass.txt")
            if err != nil {
                fmt.Println("Error reading password file:", err)
                os.Exit(1)
            }

            for _, user := range users {
                user := strings.Replace(user, "\r", "", 1)
                for _, password := range passwords {
                    pass := strings.Replace(password, "\r", "", 1)
                    // 在这里执行你的 MySQL 操作，使用当前的用户名和密码
                    if err := Plugins.MysqlScan(*singleIP, user, pass ,mysqlPort); err == nil {
                        success = true
                        break                    }
                }
                if success {
                    break
                }
            }
        }
    } else if *ipListFile != "" {
        // 从 IP 列表文件加载扫描目标
        ipList, err := readFile(*ipListFile)
        if err != nil {
            fmt.Println("Error reading IP list:", err)
            os.Exit(1)
        }

        switch mysqlType { // 修正为 mysqlType
        case "mysql":
            users, err := readFile("./dic/mysql_user.txt")
            if err != nil {
                fmt.Println("Error reading user file:", err)
                os.Exit(1)
            }

            passwords, err := readFile("./dic/mysql_pass.txt")
            if err != nil {
                fmt.Println("Error reading password file:", err)
                os.Exit(1)
            }

            for _, ip := range ipList{
                ip := strings.Replace(ip, "\r", "", 1)
                for _, user := range users {
                    user := strings.Replace(user, "\r", "", 1)
                    for _, password := range passwords {
                        pass := strings.Replace(password, "\r", "", 1)
                        // 在这里执行你的 MySQL 操作，使用当前的用户名和密码
                        if err := Plugins.MysqlScan(ip, user, pass ,mysqlPort); err == nil {
                            success = true
                            break
                        }
                    }
                    if success {
                        break
                    }
                }
            }
        }
    } else {
        fmt.Println("Usage: go run main.go")
        os.Exit(1)
    }
}

func readFile(filePath string) ([]string, error) {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    lines := strings.Split(string(content), "\n")
    return lines, nil
}
