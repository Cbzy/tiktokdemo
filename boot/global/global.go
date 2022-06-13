package global

import "gorm.io/gorm"

var (
	//DY_VP       *viper.Viper
	DyDatabase Database
	DyServer   Server
	DYDB       *gorm.DB
)

type Server struct {
	RunMode    string
	RunAddress string
	HttpPort   string
}

type Database struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    string
	MaxIdleConns uint
	MaxOpenConns uint
}
