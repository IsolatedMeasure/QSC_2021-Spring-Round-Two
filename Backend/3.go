package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID            uint
	Name          string
	Age           int
	Gender        int
	CurrentClass  string
	FreshmanClass string
	ZjuId         int64
}

func main() {
	dsn := "host=localhost user=postgres password=xusj dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed connecting")
	}

	db.Create(&Student{ID: 4, Name: "asd", Age: 1323, Gender: 0, CurrentClass: "2134123", FreshmanClass: "1acsd", ZjuId: 3200000000})
	stu := Student{}
	db.Where(&Student{Name: "小张"}).First(&stu)
	fmt.Println(stu.ZjuId)
	stus := []Student{}
	db.Where("age >= ?", "19").Find(&stus)
	for i, list := range stus {
		fmt.Println(i, ":", list.Name)
	}
	// db.First(&stu)
	db.Where(&Student{ZjuId: 3200000000}).First(&stu)
	fmt.Println("last add:", stu.Name, stu.ZjuId)
	db.Delete(&Student{}, stu.ID)
}
