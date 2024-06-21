package repository

import "gorm.io/gorm"

//go:generate go run github.com/matryer/moq -pkg mock -out ./mock/gorm_db.go . DB

type DB interface {
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}
