package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID        int64 `gorm:"primary_key"`
	Num       int64
	Bool      bool
	Str       string
	DeletedAt *time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:@/sample")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)

	// Migrate the schema
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		panic(err.Error())
	}

	// Delete
	db.Delete(&User{})

	// Create
	db.Create(&User{Num: 0, Bool: false, Str: ""})
	db.Create(&User{Num: 1, Bool: true, Str: "sample"})

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL
	db.Where(&User{ID: 0, Str: ""}).Find(&User{})

	var result User
	// recode not found
	db.Where(&User{Num: 99999999}).First(&result)

	// UPDATE `users` SET `str` = 'ALL UPDATE'  WHERE `users`.`deleted_at` IS NULL
	if err = db.Model(&result).Updates(User{Bool: false, Str: "ALL UPDATE"}).Error; err != nil {
		panic(err.Error())
	}

	// UPDATE `users` SET `deleted_at`='2024-03-24 19:51:42'  WHERE `users`.`deleted_at` IS NULL
	if err = db.Delete(&result).Error; err != nil {
		panic(err.Error())
	}
}
