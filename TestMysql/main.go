package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db = &sql.DB{}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	err = godotenv.Load()
	checkErr(err)

	db, err = sql.Open("mysql", os.Getenv("MYSQL_CONFIG"))
	checkErr(err)

	insert()
	// query()
	update()
	// query()
	delete()
}

func update() {
	//方式1 update
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		db.Exec("UPdate test set text=? where id=? ", i, i)
	}
	end := time.Now()
	fmt.Println("方式1 update total time:", end.Sub(start).Seconds())

	//方式2 update
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		stm, _ := db.Prepare("UPdate test set text=? where id=? ")
		stm.Exec(i, i)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 update total time:", end.Sub(start).Seconds())

	//方式3 update
	start = time.Now()
	stm, _ := db.Prepare("UPdate test set text=? where id=?")
	for i := 1201; i <= 1300; i++ {
		stm.Exec(i, i)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 update total time:", end.Sub(start).Seconds())

	//方式4 update
	start = time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		tx.Exec("UPdate test set text=? where id=?", i, i)
	}
	tx.Commit()

	end = time.Now()
	fmt.Println("方式4 update total time:", end.Sub(start).Seconds())

	//方式5 update
	start = time.Now()
	for i := 1401; i <= 1500; i++ {
		tx, _ := db.Begin()
		tx.Exec("UPdate test set text=? where id=?", i, i)
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 update total time:", end.Sub(start).Seconds())

}

func delete() {
	//方式1 delete
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		db.Exec("DELETE FROM test WHERE id=?", i)
	}
	end := time.Now()
	fmt.Println("方式1 delete total time:", end.Sub(start).Seconds())

	//方式2 delete
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		stm, _ := db.Prepare("DELETE FROM test WHERE id=?")
		stm.Exec(i)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 delete total time:", end.Sub(start).Seconds())

	//方式3 delete
	start = time.Now()
	stm, _ := db.Prepare("DELETE FROM test WHERE id=?")
	for i := 1201; i <= 1300; i++ {
		stm.Exec(i)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 delete total time:", end.Sub(start).Seconds())

	//方式4 delete
	start = time.Now()
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		tx.Exec("DELETE FROM test WHERE id=?", i)
	}
	tx.Commit()

	end = time.Now()
	fmt.Println("方式4 delete total time:", end.Sub(start).Seconds())

	//方式5 delete
	start = time.Now()
	for i := 1401; i <= 1500; i++ {
		tx, _ := db.Begin()
		tx.Exec("DELETE FROM test WHERE id=?", i)
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 delete total time:", end.Sub(start).Seconds())

}

func query() {

	//方式1 query
	start := time.Now()
	rows, _ := db.Query("SELECT id,user,text FROM test")
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		var content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name:%s ,id:is %d, text: %s\n", name, id, content)
	}
	end := time.Now()
	fmt.Println("方式1 query total time:", end.Sub(start).Seconds())

	//方式2 query
	start = time.Now()
	stm, _ := db.Prepare("SELECT id,user,text FROM test")
	defer stm.Close()
	rows, _ = stm.Query()
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		var content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name:%s ,id:is %d, text: %s\n", name, id, content)
	}
	end = time.Now()
	fmt.Println("方式2 query total time:", end.Sub(start).Seconds())

	//方式3 query
	start = time.Now()
	tx, _ := db.Begin()
	defer tx.Commit()
	rows, _ = tx.Query("SELECT id,user,text FROM test")
	defer rows.Close()
	for rows.Next() {
		var name string
		var id int
		var content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("name:%s ,id:is %d, text: %s\n", name, id, content)
	}
	end = time.Now()
	fmt.Println("方式3 query total time:", end.Sub(start).Seconds())
}

func insert() {

	//方式1 insert
	//strconv,int转string:strconv.Itoa(i)
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		//每次循环内部都会去连接池获取一个新的连接，效率低下
		db.Exec("INSERT INTO test(id,user,text) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
	}
	end := time.Now()
	fmt.Println("方式1 insert total time:", end.Sub(start).Seconds())

	//方式2 insert
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		//Prepare函数每次循环内部都会去连接池获取一个新的连接，效率低下
		stm, _ := db.Prepare("INSERT INTO test(id,user,text) values(?,?,?)")
		stm.Exec(i, "user"+strconv.Itoa(i), i-1000)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 insert total time:", end.Sub(start).Seconds())

	//方式3 insert
	start = time.Now()
	stm, _ := db.Prepare("INSERT INTO test(id,user,text) values(?,?,?)")
	for i := 1201; i <= 1300; i++ {
		//Exec内部并没有去获取连接，为什么效率还是低呢？ Ans:会调用dc.prepareLocked(query)
		stm.Exec(i, "user"+strconv.Itoa(i), i-1000)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 insert total time:", end.Sub(start).Seconds())

	//方式4 insert
	start = time.Now()
	//Begin函数内部会去获取连接
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		tx.Exec("INSERT INTO test(id,user,text) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
	}
	//最后释放tx内部的连接
	tx.Commit()

	end = time.Now()
	fmt.Println("方式4 insert total time:", end.Sub(start).Seconds())

	//方式5 insert
	start = time.Now()
	for i := 1401; i <= 1500; i++ {
		//Begin函数每次循环内部都会去连接池获取一个新的连接，效率低下
		tx, _ := db.Begin()
		tx.Exec("INSERT INTO test(id,user,text) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
		//Commit执行后连接也释放了
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 insert total time:", end.Sub(start).Seconds())
}
