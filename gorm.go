package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func create(db *gorm.DB) {
	//创建表
	db.AutoMigrate(&Product{})
}

func insertData(db *gorm.DB) {
	//插入数据
	p := Product{
		Code:  "1001",
		Price: 20,
	}
	db.Create(&p)
}

//查找
func find(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	fmt.Printf("p: %v\n", p)
	db.First(&p, "code=?", "1001")
	fmt.Printf("p: %v\n", p)
}

//更新
func update(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Model(&p).Update("price", 30)
	//db.Model(&p).Updates(Product{Price: 33, Code: "1002"})

	db.Model(&p).Updates(map[string]interface{}{"Price": 200, "Code": "1003"})
}

//删除
func de(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Delete(&p, 1)
}

/* func main() {
	dns := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		//fmt.Printf("err: %v\n", err)
		panic("failed to connect database")
	}
	//find(db)
	//update(db)
	de(db)
}
*/
