package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dd *gorm.DB

func init() {
	dns := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	d1, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		//fmt.Printf("err: %v\n", err)
		panic("failed to connect database")
	}
	dd = d1
}

func test2() {
	type Company struct {
		ID   int
		Name string
	}

	type User2 struct {
		gorm.Model
		Name      string
		CompanyID int
		Company   Company
	}

	dd.AutoMigrate(&Company{}, &User2{})
}

func test3() {
	type CreditCard struct {
		gorm.Model
		Number  string
		User3ID uint
	}

	// User 有一张 CreditCard，UserID 是外键
	type User3 struct {
		gorm.Model
		CreditCard CreditCard
	}
	dd.AutoMigrate(&User3{}, &CreditCard{})

}

func test4() {
	type Toy struct {
		ID        int
		Name      string
		OwnerID   int
		OwnerType string
	}

	type Cat struct {
		ID   int
		Name string
		Toy  Toy `gorm:"polymorphic:Owner;"`
	}

	type Dog struct {
		ID   int
		Name string
		Toy  Toy `gorm:"polymorphic:Owner;"`
	}
	dd.AutoMigrate(&Toy{}, &Cat{}, &Dog{})
	dd.Create(&Dog{Name: "dog", Toy: Toy{Name: "toy1"}})

}

func test5() {
	//Has Many
	type CreditCard struct {
		gorm.Model
		Number  string
		User4ID uint
	}
	// User 有多张 CreditCard，User4ID 是外键
	type User4 struct {
		gorm.Model
		CreditCards []CreditCard
	}
	dd.AutoMigrate(&User4{}, &CreditCard{})
}
func main() {
	test5()
}
