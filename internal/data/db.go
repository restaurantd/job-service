package data

import "gorm.io/gorm"

type DataBase struct {
	*gorm.DB
}
