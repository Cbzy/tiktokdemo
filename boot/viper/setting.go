package viper

import (
	"douyin/boot/global"
	"fmt"
	"github.com/spf13/viper"
)

type Viper struct {
	vp *viper.Viper
}

func NewSetting() (*Viper, error) {
	vp := viper.New()

	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	vp.AddConfigPath("config")
	vp.AddConfigPath("config.yaml")
	err := vp.ReadInConfig()

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	global.DyServer = global.Server{
		RunMode:    vp.GetString("Server.RunMode"),
		RunAddress: vp.GetString("Server.RunAddress"),
		HttpPort:   vp.GetString("Server.HttpPort"),
	}

	//if err := vp.Unmarshal(&global.DyDatabase); err != nil {
	//	fmt.Println(err)
	//}

	global.DyDatabase = global.Database{
		DBType:       vp.GetString("Database.DBType"),
		Username:     vp.GetString("Database.Username"),
		Password:     vp.GetString("Database.Password"),
		Host:         vp.GetString("Database.Host"),
		DBName:       vp.GetString("Database.DBName"),
		TablePrefix:  vp.GetString("Database.TablePrefix"),
		Charset:      vp.GetString("Database.Charset"),
		ParseTime:    vp.GetString("Database.ParseTime"),
		MaxIdleConns: uint(vp.GetInt("Database.MaxIdleConns")),
		MaxOpenConns: uint(vp.GetInt("Database.MaxOpenConns")),
	}

	return &Viper{vp}, nil
}
