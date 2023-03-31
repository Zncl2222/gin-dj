package app

import (
	"github.com/Zncl2222/gin-dj/core"

	"gorm.io/gorm"
)

type (
	ExampleModel struct {
		gorm.Model
		Content string `json:"content"`
	}
)

func AutoMigrate() {
	db := core.GetDB()
	db.AutoMigrate(&ExampleModel{})
}
