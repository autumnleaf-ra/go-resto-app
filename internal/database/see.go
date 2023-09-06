package database

import (
	"github.com/autumnleaf-ra/go-resto-app/internal/model"
	"github.com/autumnleaf-ra/go-resto-app/internal/model/constant"
	"gorm.io/gorm"
)

func SeeDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica",
			Price:     41250,
			Type:      constant.MenuTypeFood,
		},
	}

	drinksMenu := []model.MenuItem{
		{
			Name:      "Lemon Tea",
			OrderCode: "lemon_tea",
			Price:     6000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Jus Jeruk",
			OrderCode: "jeruk",
			Price:     7500,
			Type:      constant.MenuTypeDrink,
		},
	}

	db.AutoMigrate(&model.MenuItem{})

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}
}
