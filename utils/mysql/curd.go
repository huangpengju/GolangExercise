package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	age  int
	name string
}

// 初始化
func InitDB() (err error) {
	dsn := "root:Aa_123456@tcp(127.0.1.1:3306)/sql_test"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		// panic(err)
		return err
	}
	err = db.Ping()

	if err != nil {
		return err
	}
	return nil
}

// 查询单条数据展示
func QueryRowDemo() {

	sqlStr := "select id,name,age from user where id=?"
	var u user
	//非常重要：确保QuerRow之后调用Scan方法，否则

	err := db.QueryRow(sqlStr, 6).Scan(&u.id, &u.name, &u.age)

	if err != nil {
		fmt.Printf("Scan失败%v\n", err)
		return
	}

	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 多条查询
func QueryMyltiRowDemo() {
	sqlStr := "select id,name,age from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("请求失败err%v\n", err)
		return
	}
	defer rows.Close()
	//循环父亲结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("Scan失败，err%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func InsertRowDemo() {
	sqlStr := "insert into user(name,age) values(?,?)"
	ret, err := db.Exec(sqlStr, "王五", 22)
	if err != nil {
		fmt.Printf("insert失败,err%v\n", err)
		return
	}
	theId, err := ret.LastInsertId() //新插入数据的id
	if err != nil {
		fmt.Printf("get 新插入数据id失败,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d.\n", theId)
}

// 更新数据
func UpdateRowDemo() {
	sqlStr := "Update user set age=? where id=?"
	ret, err := db.Exec(sqlStr, 32, 7)
	if err != nil {
		fmt.Printf("update 失败，err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected失败，err:%v\n", err)
		return
	}
	fmt.Printf("update success,affected rows:%d\n", n)
}

// 删除数据
func DeleteRowDemo() {
	sqlStr := "delete from user where id=?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("delete 失败，err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected 失败，err:%v\n", err)
		return
	}
	fmt.Printf("delete success ,affected rows:%d\n", n)
}
