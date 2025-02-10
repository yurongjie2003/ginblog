package config

import (
	"errors"
	"gopkg.in/ini.v1"
)

var (
	AppMode          string
	HttpPort         string
	DbHost           string
	DbPort           string
	DbUser           string
	DbPassword       string
	DbName           string
	JwtSecret        string
	JwtEffectiveTime int64
	JwtIssuer        string
)

func Init() error {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		return errors.New("the configuration file failed to load correctly")
	}
	loadServer(file)
	loadDatabase(file)
	loadJwt(file)
	return nil
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8080")
}

func loadDatabase(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}

func loadJwt(file *ini.File) {
	JwtSecret = file.Section("jwt").Key("JwtSecret").MustString("default_secret")
	JwtEffectiveTime = file.Section("jwt").Key("JwtEffectiveTime").MustInt64(360)
	JwtIssuer = file.Section("jwt").Key("JwtIssuer").MustString("ginblog")
}
