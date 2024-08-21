package oss

import (
	"context"
	"io/ioutil"
)

func (c *OSSClient) DownloadStream(ctx context.Context, bucketName string, key string) ([]byte, error) {
	bucket, err := c.getBucket(bucketName)
	if err != nil {
		return nil, err
	}
	body, err := bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
