package Minio

import (
	"github.com/yurongjie2003/ginblog/utils/Config"
)

var Client *MinIOClient

func Init() error {
	var err error
	Client, err = NewMinIOClient(
		Config.MinioEndpoint,
		Config.MinioAccessKey,
		Config.MinioSecretKey,
		Config.MinioBucketName,
		Config.MinioLocation,
		Config.MinioPrivateFolderPath,
		Config.MinioPublicFolderPath,
	)
	return err
}
