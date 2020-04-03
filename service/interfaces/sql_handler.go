package interfaces

import (
	"github.com/jinzhu/gorm"
)

// SQLHandler interface
type SQLHandler interface {
	Where(dest interface{}, query string, args interface{}) *gorm.DB
	Create(dest interface{}) *gorm.DB
	Delete(dest interface{}, query string, args ...interface{}) *gorm.DB
}
