package model

import (
	utils "GolangExercise/utils/mySQL"
	"fmt"
)

func LinkMysql() {
	err := utils.InitDB()

	if err != nil {

		fmt.Printf("初始化失败%v\n", err)

		return
	}

	//单行查询
	utils.QueryRowDemo()

	//插入数据
	utils.InsertRowDemo()
	//更新数据
	utils.UpdateRowDemo()
	//删除数据
	utils.DeleteRowDemo()
	//多行查询
	utils.QueryMyltiRowDemo()

}
