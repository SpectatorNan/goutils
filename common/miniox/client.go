package miniox

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"mime/multipart"
)

type Config struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	Location string
}

type Client struct {
	MinioClient *minio.Client
	location    string
}

func NewClient(config Config) *Client {
	minioClient, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil
	}
	return &Client{
		MinioClient: minioClient,
		location:    config.Location,
	}
}

func (c *Client) CreateBucket(ctx context.Context, bucketName string) error {
	// Check to see if we already own this bucket (which happens if you run this twice)
	exists, errBucketExists := c.MinioClient.BucketExists(ctx, bucketName)
	if errBucketExists != nil {
		return errBucketExists
	}
	if exists {
		return nil
	}

	err := c.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: c.location})
	if err != nil {
			return err
	}
		return nil

}

func (c *Client) UploadFile(ctx context.Context, bucketName, objectName string, file multipart.File, fileStat *multipart.FileHeader) error {

	uploadInfo, err := c.MinioClient.PutObject(ctx, bucketName, objectName, file, fileStat.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
	return nil
}