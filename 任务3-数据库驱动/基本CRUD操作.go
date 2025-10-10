package main

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql" // 或其他数据库驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	username := "root"
	password := "123456"
	host := "localhost"
	port := "3306"
	dabase := "testdb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dabase)

	// mysql打印日志
	mysqlLogger := logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})

	if err != nil {
		log.Fatalf("failed to connect to the database : %v", err)
	}

	fmt.Println("database connect success!!!!!!!!!!")

	DB = db
}

type Student struct {
	Id    uint   `gorm:"primarykey;autoIncrement"`
	Name  string `gorm:"size:64"`
	Age   uint8
	Grade string `gorm:"size:64"`
}

func main() {
	ctx := context.Background()

	stu := Student{Name: "张三", Age: 18, Grade: "三年级"}

	DB.AutoMigrate(&stu)
	DB.Save(&stu)

	stru2, err := gorm.G[Student](DB).Where("age > ", 18).Find(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stru2)

	gorm.G[Student](DB).Where("name", "张三").Update(ctx, "grade", "四年级")

	gorm.G[Student](DB).Where("age <= ?", 15).Delete(ctx)

}
