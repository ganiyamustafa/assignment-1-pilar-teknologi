package utils

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Handler struct {
	SQLite    *gorm.DB
	Validator *validator.Validate
}
