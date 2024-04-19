package types

type AppConfig struct {
	PRODUCTION                        bool
	PUBLIC_HOST                       string
	PORT                              string
	DB_USER                           string
	DB_PASSWORD                       string
	DB_HOST                           string
	DB_PORT                           string
	DB_NAME                           string
	JWT_ACCESS_SECRET                 string
	JWT_REFRESH_SECRET                string
	JWT_ACCESS_EXPIRATION_IN_SECONDS  int64
	JWT_REFRESH_EXPIRATION_IN_SECONDS int64
}
