package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db1 *gorm.DB

func init() {
	dns := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	d1, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		//fmt.Printf("err: %v\n", err)
		panic("failed to connect database")
	}
	db1 = d1
}

type User1 struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func creates() {
	db1.AutoMigrate(&User1{})
}

var us = User1{
	Name:     "wom",
	Age:      22,
	Birthday: time.Now(),
}

func insert1() {

	d := db1.Create(&us)
	fmt.Printf("d.RowsAffected: %v\n", d.RowsAffected)
	fmt.Printf("us.ID: %v\n", us.ID)
}

func insert2() {
	db1.Select("Name", "Age", "CreateAt").Create(&us)
}

func insert3() {
	var p = []User1{{Name: "lili"}, {Name: "zhangsan"}, {Name: "wangwu"}}
	db1.Create(&p)
}

func select1() {
	var u User1
	//db1.First(&u)
	// fmt.Printf("u.ID: %v\n", u.ID)
	// db1.Take(&u)
	// fmt.Printf("u.ID: %v\n", u.ID)
	db1.Last(&u)
	fmt.Printf("u.ID: %v\n", u.ID)
}

type Result struct {
	Id   int
	Name string
	Age  int
}

func testRaw() {
	var r Result
	db1.Raw("select id, name, age from user1 where name=? ", "tom").Scan(&r)
	fmt.Printf("r: %v\n", r)
	var ag int
	db1.Raw("select sum(age) from user1").Scan(&ag)
	fmt.Printf("ag: %v\n", ag)
}

func testRaw2() {
	db1.Exec("update user1 set name = ? where id = ?", "lili", 3)
}

func testRaw3() {
	var user User1
	db1.Where("name = @myname", sql.Named("myname", "lili")).Find(&user)
	fmt.Printf("user: %v\n", user)
}

func testRaw4() {
	var name string
	var age int
	// 使用 GORM API 构建 SQL
	row := db1.Table("user1").Where("name = ?", "wom").Select("name", "age").Row()
	row.Scan(&name, &age)
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
}

/* func main() {
	testRaw4()
} */
