package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetAuthAdmin() (string, string) {
	viper.SetConfigFile("./config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	user := viper.GetString("AuthUser.Admin0.name")
	passwd := viper.GetString("AuthUser.Admin0.passwd")
	return user, passwd
}
