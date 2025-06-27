package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ProductId uint
	qty       int
	Total     int
}
