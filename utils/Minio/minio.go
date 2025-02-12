package Minio

// go get github.com/minio/minio-go/v7

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/url"
	"time"
)

type MinIOClient struct {
	client            *minio.Client
	bucketName        string
	location          string
	PrivateFolderPath string // 例如 "private-folder/"
	PublicFolderPath  string // 例如 "public-folder/"
}

// NewMinIOClient 初始化MinIO客户端
func NewMinIOClient(endpoint, accessKey, secretKey, bucketName, location, privateFolderPath, publicFolderPath string) (*MinIOClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: false, // 根据实际情况调整
	})
	if err != nil {
		return nil, err
	}

	mc := &MinIOClient{
		client:            client,
		bucketName:        bucketName,
		location:          location,
		PrivateFolderPath: privateFolderPath,
		PublicFolderPath:  publicFolderPath,
	}

	// 创建存储桶（如果不存在）
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return nil, err
		}
	}

	// 配置存储桶策略
	if err := mc.SetReadOnlyPolicy(); err != nil {
		return nil, fmt.Errorf("failed to set bucket policy: %v", err)
	}

	return mc, nil
}

// SetReadOnlyPolicy 设置只读存储桶策略
func (m *MinIOClient) SetReadOnlyPolicy() error {
	policy := fmt.Sprintf(`{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {"AWS": ["*"]},
				"Action": ["s3:GetObject"],
				"Resource": ["arn:aws:s3:::%s/%s*"]
			}
		]
	}`, m.bucketName, m.PublicFolderPath)

	ctx := context.Background()
	err := m.client.SetBucketPolicy(ctx, m.bucketName, policy)
	return err
}

// UploadFile 上传文件到指定路径
func (m *MinIOClient) UploadFile(file io.Reader, objectPath string, fileSize int64, contentType string) error {
	ctx := context.Background()

	_, err := m.client.PutObject(
		ctx,
		m.bucketName,
		objectPath,
		file,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	return err
}

// UploadPublicFile 上传公共文件
func (m *MinIOClient) UploadPublicFile(file io.Reader, objectName string, fileSize int64, contentType string) error {
	objectPath := m.PublicFolderPath + objectName
	return m.UploadFile(file, objectPath, fileSize, contentType)
}

// GetPublicFileURL 获取公共文件访问URL
func (m *MinIOClient) GetPublicFileURL(objectName string) string {
	return fmt.Sprintf("%s/%s/%s%s", m.client.EndpointURL(), m.bucketName, m.PublicFolderPath, objectName)
}

// UploadPrivateFile 上传私有文件
func (m *MinIOClient) UploadPrivateFile(file io.Reader, objectName string, fileSize int64, contentType string) error {
	objectPath := m.PrivateFolderPath + objectName
	return m.UploadFile(file, objectPath, fileSize, contentType)
}

// GeneratePrivateFileURL 获取私有文件访问URL(默认7天有效期)
func (m *MinIOClient) GeneratePrivateFileURL(objectName string, expires time.Duration) (string, error) {
	if expires == 0 {
		expires = time.Hour * 24 * 7 // 默认7天
	}

	ctx := context.Background()
	objectPath := m.PrivateFolderPath + objectName

	reqParams := make(url.Values)
	privateFileURL, err := m.client.PresignedGetObject(
		ctx,
		m.bucketName,
		objectPath,
		expires,
		reqParams,
	)
	if err != nil {
		return "", err
	}

	return privateFileURL.String(), nil
}
