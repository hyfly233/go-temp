package global

import "gorm.io/gorm"

var (
	ENV *string
	DB  *gorm.DB
)
