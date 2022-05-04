package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *sql.DB //连接池对象

type user struct { //结构体
	id       int
	username string
	password string
}

func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:@tcp(127.0.0.1:3306)/test"
	db, err = sql.Open("mysql", dsn) //open不会校验用户名和密码是否正确
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10) //设置数据库连接池的最大连接数 10
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

func queryOne(id int) {
	var u1 user
	// 1.写查询单条记录的sql语句
	sqlStr := "select id,username,password from user where id =?;"
	//2.执行并拿到结果
	//必须对row对象调用scan方法，该方法会释放数据库连接
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.username, &u1.password) //从连接池拿一个连接出去数据库查询单条记录
	//打印结果
	fmt.Println(u1)
}
func queryMore(n int) {
	//1.sql语句
	sqlStr := "select id,username,password from user where id >?;"
	//2.执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println(err)
		return
	}
	//3.一定要关闭rows
	defer rows.Close()
	//4.循环取值
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.username, &u1.password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(u1)
	}
}
func insert() {
	//写sql语句
	sqlStr := "insert into user(username,password) values ('feifei','ssssss')"
	//exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed %v\n", err)
		return
	}
	//如果是插入操作，可以拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed %v\n", err)
		return
	}
	fmt.Println("id:", id)
}
func updateRow() {
	sqlStr := "update user set password='passwd' where id > 3"
	ret, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Printf("update failed %v\n", err)
		return
	}
	//RowsAffected获取更新了几行
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get n failed %v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据", n)
}

func upd() {
	s := "update user set username=?,password=? where id = ?"
	r, err := db.Exec(s, "lili", "23452", "1")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}
}

func del() {
	s := "delete from user where id=?"
	r, err := db.Exec(s, 1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		i, _ := r.RowsAffected()
		fmt.Printf("i: %v\n", i)
	}
}
func deleteRow(id int) {
	sqlStr := "delete from user where id >?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get n failed %v\n", err)
		return
	}
	fmt.Printf("删除了%d行数据", n)
}

var client *mongo.Client

func initmongo() {
	clientOpt := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("client: %v\n", client)
	fmt.Println("Mongodb连接成功！")

}

type Student struct {
	Name string
	Age  int
}

func insertmongo() {
	s := Student{
		Name: "lili",
		Age:  20,
	}
	// c2 := client.Database("go_db").Collection("stu")
	c := client.Database("go_db").Collection("user")
	ior, err := c.InsertOne(context.TODO(), s)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("ior.InsertedID: %v\n", ior.InsertedID)
	}

}

func inserMany() {
	initDB()
}

/* func main() {

	d := bson.D{{"name", "tom"}}
	fmt.Printf("d: %v\n", d)

	// initmongo()
	// insertmongo()
	// err := initDB()
	// if err != nil {
	// 	fmt.Printf("init DB failed err:%v\n", err)
	// }
	// fmt.Println("连接数据库成功")

	// queryOne(2)
	// fmt.Println("------")
	// queryMore(0)
	// fmt.Println("------")
	// insert()
	// upd()
	// del()
	// deleteRow(4)
}
*/
