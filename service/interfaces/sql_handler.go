package interfaces

// SQLHandler interface
type SQLHandler interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Create(dest interface{}) error
	Update(model interface{}, column string, newData interface{}, query string, args ...interface{}) error
}
