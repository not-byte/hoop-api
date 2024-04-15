package store

import (
	"tournament_api/server/types"

	_ "github.com/lib/pq"
)

type SQLStore struct{}

func NewSQLStore(config *types.AppConfig) *SQLStore {
	/*

		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		db, err := sql.Open("postgres", psqlInfo)

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully connected to CockroachDB!")
	*/

	return &SQLStore{}
}

func (s *SQLStore) Get() any {
	var value any = "mock"
	return value
}
