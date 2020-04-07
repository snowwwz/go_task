package infrastructure

import (
	"github.com/davecgh/go-spew/spew"
	// DB driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// SQLHandler DB connection
type SQLHandler struct {
	Conn *gorm.DB
}

// Connect get SQLHandler
func Connect() *SQLHandler {

	db, err := gorm.Open("mysql", "yukino:aaa@tcp(localhost:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		spew.Dump(err.Error())
	}

	// 疎通確認

	return &SQLHandler{
		Conn: db,
	}
}

// Select with where
func (s *SQLHandler) Select(dest interface{}, query string, args ...interface{}) error {
	return s.Conn.Where(query, args...).Find(dest).Error
}

// Create insert
func (s *SQLHandler) Create(dest interface{}) error {
	return s.Conn.Create(dest).Error
}

// Update update
func (s *SQLHandler) Update(model interface{}, column string, newData interface{}, query string, args ...interface{}) error {
	return s.Conn.Model(model).Where(query, args...).Update(column, newData).Error
}
