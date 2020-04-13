package infrastructure

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	// DB driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yukinooz/go_task/config"
)

// SQLHandler DB connection
type SQLHandler struct {
	Conn *gorm.DB
}

// Connect get SQLHandler
func Connect(conf config.DBconf) *SQLHandler {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/todo?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Pass, conf.Host, conf.Port)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		spew.Dump(err.Error())
	}

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
