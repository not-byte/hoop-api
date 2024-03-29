package store

type SQLStore struct{}

func NewSQLStore() *SQLStore {
	return &SQLStore{}
}
func (s *SQLStore) Get() any {
	var value any = "mock"
	return value
}
