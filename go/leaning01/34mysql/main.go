package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id int
	age int
	name string
}

// 连接数据库
func connectDB()(err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// 查询单条数据
func query(id int) (u user, v error) {
	sqlStr := "select id, name, age from user where id=?"

	// 非常重要: 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	if err == sql.ErrNoRows {
		fmt.Printf("user data that does not exist!")
		return u, err
	}
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return u, err
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)

	return u, nil
}

// 多行查询
func queryMultiRow(where string, fields string) (users []user, err error)  {
	sqlStr := "select  " + fields + " from user where " + where
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return nil, err
	}
	// 非常重要:关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, err
		}

		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
		users = append(users, u)
	}

	return users, nil
}


// 插入数据
func insert(name string, age int) (id int64, err error) {
	sqlStr := "insert into user(name,age) values(?, ?)"
	ret, err := db.Exec(sqlStr, name, age)

	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return 0, err
	}

	ID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return 0, err
	}
	fmt.Printf("Insert success, the id is %d.\n", ID)

	return ID, nil
}


// 更新数据
func update(update string, where string) (updatedRows int64, err error) {
	sqlStr := "update user set " + update + " where " + where
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return 0, err
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get rowsAffected failed, err:%v\n", err)
		return 0, err
	}

	fmt.Printf("update success, affected rows:%d\n", n)
	return n, nil
}


// 删除数据
func delete(where string) (affectedRows int64, err error) {
	sqlStr := "delete from user where " + where
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return 0, err
	}

	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed, err:%v\n", err)
		return 0, err
	}

	fmt.Printf("delete suceess, affected rows:%d\n", n)
	return n, nil
}


func main() {
	err := connectDB()
	if err != nil {
		fmt.Printf("connect db fail, err:%v\n", err)
		return
	}

	// fmt.Printf("%#v\n", db)

	// insert("Jack", 18)
	// insert("Tom", 20)

	// update("age=30", "id=2")

	// delete("id=2")
}
