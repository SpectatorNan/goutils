package oss

import (
	"context"
	"io"
	"mime/multipart"
)

func (c *OSSClient) UploadStreamByFile(ctx context.Context, bucketName string, key string, file multipart.File) error {
	bucket, err := c.getBucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.PutObject(key, file)
	return err
}

func (c *OSSClient) UploadStream(ctx context.Context, bucketName string, key string, stream io.Reader) error {
	bucket, err := c.getBucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.PutObject(key, stream)
	return err
}
