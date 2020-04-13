package config

type Config struct {
	DB  DBconf
	Log Logconf
}

type DBconf struct {
	Host string
	Port string
	User string
	Pass string
}

type Logconf struct {
	File string
}
