package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

// 1.读取yarn配置
type YarnMysqlConfig struct {
	Address  string
	Port     int
	Username string
	Password string
}

func main() {
	var mysqlConfig IniMysqlConfig

	conf, err := ini.Load(&mysqlConfig, "D:\\godemo\\goBase\\day5\\myconfig.ini")

	if err != nil {
		fmt.Println("读取配置文件失败：", err)
		return
	}
	if err := conf.Section("mysql").MapTo(&mysqlConfig); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("读取配置文件成功：")
	fmt.Println("主机：", mysqlConfig.Address)
	fmt.Println("端口：", mysqlConfig.Port)
	fmt.Println("用户名：", mysqlConfig.Username)
	fmt.Println("密码：", mysqlConfig.Password)
}
