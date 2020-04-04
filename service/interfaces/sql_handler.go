package interfaces

// SQLHandler interface
type SQLHandler interface {
	SelectAll(dest interface{}, query string, args ...interface{}) error
	Create(dest interface{}) error
	Delete(model interface{}, query string, args ...interface{}) error
}
