package Config

import (
	"errors"
	"gopkg.in/ini.v1"
)

var (
	AppMode                string
	HttpPort               string
	MaxFileSize            int64
	DbHost                 string
	DbPort                 string
	DbUser                 string
	DbPassword             string
	DbName                 string
	JwtSecret              string
	JwtEffectiveTime       int64
	JwtIssuer              string
	MinioEndpoint          string
	MinioAccessKey         string
	MinioSecretKey         string
	MinioBucketName        string
	MinioLocation          string
	MinioPrivateFolderPath string
	MinioPublicFolderPath  string
)

func Init() error {
	file, err := ini.Load("utils/Config/Config.ini")
	if err != nil {
		return errors.New("the configuration file failed to load correctly")
	}
	loadServer(file)
	loadDatabase(file)
	loadJwt(file)
	loadMinio(file) // 新增
	return nil
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8080")
	MaxFileSize = file.Section("server").Key("MaxFileSize").MustInt64(5242880)
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

func loadMinio(file *ini.File) { // 新增
	MinioEndpoint = file.Section("minio").Key("endpoint").MustString("")
	MinioAccessKey = file.Section("minio").Key("accessKey").MustString("")
	MinioSecretKey = file.Section("minio").Key("secretKey").MustString("")
	MinioBucketName = file.Section("minio").Key("bucketName").MustString("")
	MinioLocation = file.Section("minio").Key("location").MustString("")
	MinioPrivateFolderPath = file.Section("minio").Key("privateFolderPath").MustString("")
	MinioPublicFolderPath = file.Section("minio").Key("publicFolderPath").MustString("")
}
