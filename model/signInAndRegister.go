package model

import (
	utils "GolangExercise/utils/signInAndRegister"
	"fmt"
)

func SingInAndRegister() {
	// 定义密码
	password := "123456"
	// 尝试第一次加密密码
	encrypt1, _ := utils.EncryptPassword(password)
	// 打印第一次加密的密码
	fmt.Printf("第一次加密密码：%v\n", encrypt1)
	// 尝试第二次加密密码
	encrypt2, _ := utils.EncryptPassword(password)
	// 打印第二次加密的密码
	fmt.Printf("第一次加密密码：%v\n", encrypt2) // 可以发现两次加密密码是不一样的结果

	// 对比密码是否正确，第一次加密的字符串加密对比
	passwordEquals := utils.EqualsPassword(password, encrypt1)
	fmt.Printf("使用第一次加密的密码字符串对比密码是否正确：%v\n", passwordEquals)

	// 对比密码是否正确，第二次加密的字符串加密对比
	passwordEquals = utils.EqualsPassword(password, encrypt2)
	fmt.Printf("使用第二次加密的密码字符串对比密码是否正确：%v\n", passwordEquals)

	// 尝试对比一个错误的密码
	passwordEquals = utils.EqualsPassword("1234567", encrypt1) // 1234567 是错误的密码
	fmt.Printf("尝试用错误的密码对比密码是否正确：%v\n", passwordEquals)
}
