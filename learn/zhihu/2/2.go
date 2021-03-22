package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//自动检查 Product 结构是否变化，变化则进行迁移
	db.AutoMigrate(&Animal{})

	var animal = Animal{Age: 99, Name: ""}
	db.Create(&animal)

	fmt.Println(db.First(&animal))
	//// 增
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//// 查
	//var product Product
	//db.First(&product, 1) // 找到id为1的产品
	//db.First(&product, "code = ?", "L1212") // 找出 code 为 l1212 的产品
	//
	//// 改 - 更新产品的价格为 2000
	//db.Model(&product).Update("Price", 2000)

	// 删 - 删除产品
	//db.Delete(&product)
}
