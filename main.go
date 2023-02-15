package main

import (
	"GolangExercise/model"
	"fmt"
)

func main() {
	//1.连接数据库
	fmt.Println("————————1.连接数据库————————1")
	model.LinkMysql()
	fmt.Println()

	fmt.Println("————————2.加密和验证密码连接数据库————————1")
	//2.加密和验证密码
	model.SingInAndRegister()
}
