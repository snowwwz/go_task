package infrastructure

import "github.com/jinzhu/gorm"

// SQLHandler struct
type SQLHandler struct {
	Conn *gorm.DB
}

// NewHandler create a new handler
func NewHandler() (*SQLHandler, error) {
	conn, err := gorm.Open("mysql", "yukino:aaa@tcp(localhost:3306)/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	// if err := conn.; err != nil {
	// 	return nil, err
	// }疎通確認

	return &SQLHandler{
		Conn: conn,
	}, nil
}
