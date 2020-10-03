package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id int
	name string
	age int8
}

func connectDB()(err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("cennect sql failed. err:", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("cennect sql failed. err:", err)
		return err
	}

	return nil
}


// 事务操作示例
func transaction() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Println("begin transaction failed, err:%v\n", err)
		return
	}

	sqlStr1 := "Update user set age=30 where id=?"
	ret1, err := db.Exec(sqlStr1, 1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err: %v\n", err)
		return
	}

	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "insert user(name, age) value(?, ?)"
	ret2, err := db.Exec(sqlStr2, "Tom", 24)
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}

	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback()
		fmt.Printf("exec ret2.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("提交事务。")
		tx.Commit()
	} else {
		tx.Rollback()
		fmt.Println("事务回滚...")
	}

	fmt.Println("exec trans success!")
}


func main() {
	err := connectDB()	
	if err != nil {
		fmt.Println("database connect failed. err:", err)
		return 
	}

	transaction()
}
