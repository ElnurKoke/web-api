package model

type Config struct {
	Port string
	DB   struct {
		Dsn    string
		Driver string
	}
}
