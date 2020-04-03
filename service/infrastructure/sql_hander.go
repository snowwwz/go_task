package infrastructure

import (
	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func Connect() *SQLHandler {

	db, err := gorm.Open("mysql", "yukino:aaa@tcp(localhost:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		spew.Dump(err.Error())
	}

	// if err := conn.; err != nil {
	// 	return nil, err
	// }疎通確認

	return &SQLHandler{
		Conn: db,
	}
}

// Where select
func (s *SQLHandler) Where(dest interface{}, query string, args interface{}) *gorm.DB {
	return s.Conn.Find(dest, query, args)
}

// Create insert
func (s *SQLHandler) Create(dest interface{}) *gorm.DB {
	return s.Conn.Create(dest)
}

// Delete delete
func (s *SQLHandler) Delete(dest interface{}, query string, args ...interface{}) *gorm.DB {
	return s.Conn.Model(dest).Where(query, args...).Update("delete_flg", 1)
}
