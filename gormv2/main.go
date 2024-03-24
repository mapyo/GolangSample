package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int64 `gorm:"primary_key"`
	Num       int64
	Bool      bool
	Str       string
	DeletedAt gorm.DeletedAt
}

func main() {
	dsn := "root:@/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()

	// Migrate the schema
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	// WHERE conditions required
	db.Delete(&User{})

	// Create
	db.Create(&User{Num: 0, Bool: false, Str: ""})
	db.Create(&User{Num: 1, Bool: true, Str: "sample"})

	// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL
	db.Where(&User{ID: 0, Str: ""}).Find(&User{})

	var result User
	// recode not found
	// SELECT * FROM `users` WHERE `users`.`num` = 99999999 AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1
	db.Where(&User{Num: 99999999}).First(&result)

	// WHERE conditions required
	// UPDATE `users` SET `str`='ALL UPDATE' WHERE `users`.`deleted_at` IS NULL
	db.Model(&result).Updates(User{Bool: false, Str: "ALL UPDATE"})

	// WHERE conditions required
	// UPDATE `users` SET `deleted_at`='2024-03-24 19:58:13.514' WHERE `users`.`deleted_at` IS NULL
	db.Delete(&result)
}
