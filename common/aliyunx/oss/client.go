package oss

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type Config struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
}

type OSSClient struct {
	client *oss.Client
}

func NewOSSClient(cfg Config) *OSSClient {
	client, err := oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		panic(err)
	}
	return &OSSClient{client: client}
}

func (c *OSSClient) getBucket(name string) (*oss.Bucket, error) {
	bucket, err := c.client.Bucket(name)
	if err != nil {
		return nil, err
	}
	return bucket, err
}
