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

type Accounts struct {
	ID      uint    `gorm:"primarykey;auto_increment" json:"id"`
	Balance float64 `gorm:"type:decimal(10, 2);not null;default:0"`
}

type Transactions struct {
	ID             uint    `gorm:"primarykey;auto_increment" json:"id"`
	From_acount_id uint    `gorm:"not null" json:"from_acount_id"`
	To_acount_id   uint    `gorm:"not null" json:"to_acount_id"`
	Amount         float64 `gorm:"type:decimal(10, 2);not null;default:0"`
}

func main() {
	ctx := context.Background()

	a := Accounts{Balance: 10.0}
	b := Accounts{Balance: 1.0}

	DB.AutoMigrate(&a)
	DB.AutoMigrate(&b)
	tr := Transactions{}
	DB.AutoMigrate(tr)

	DB.Create(&a)
	DB.Create(&b)

	tx := DB.Begin()
	if a.Balance > 50 {
		a.Balance -= 50
		b.Balance += 50
	} else {
		fmt.Println("A 账户余额不足！！！")
		tx.Rollback()
	}
	if err := gorm.G[Transactions](DB).Create(ctx, &Transactions{From_acount_id: a.ID, To_acount_id: b.ID, Amount: 50}); err != nil {
		fmt.Println("交易记录失败：", err)
		tx.Rollback()
	}

	tx.Commit()
}
