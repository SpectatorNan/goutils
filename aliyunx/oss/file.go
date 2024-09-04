package oss

import "context"

func (c *OSSClient) DeleteObject(ctx context.Context, bucketName string, key string) error {
	bucket, err := c.getBucket(bucketName)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(key)
	return err
}

func (c *OSSClient) DeleteObjects(ctx context.Context, bucketName string, keys []string) error {
	if len(keys) == 0 {
		return nil
	}
	if len(keys) < 1000 {
		_, err := c.batchDeleteObjects(ctx, bucketName, keys)
		return err
	}
	for i := 0; i < len(keys); i += 1000 {
		end := i + 1000
		if end > len(keys) {
			end = len(keys)
		}
		_, err := c.batchDeleteObjects(ctx, bucketName, keys[i:end])
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *OSSClient) DeleteObjectsWithFailed(ctx context.Context, bucketName string, keys []string) ([]string, error) {
	if len(keys) == 0 {
		return nil, nil
	}
	if len(keys) < 1000 {
		failedKeys, err := c.batchDeleteObjects(ctx, bucketName, keys)
		return failedKeys, err
	}
	var failedKeys []string
	for i := 0; i < len(keys); i += 1000 {
		end := i + 1000
		if end > len(keys) {
			end = len(keys)
		}
		res, err := c.batchDeleteObjects(ctx, bucketName, keys[i:end])
		if err != nil {
			failedKeys = append(failedKeys, res...)
		}
	}
	return failedKeys, nil
}

// keys max length is 1000
func (c *OSSClient) batchDeleteObjects(ctx context.Context, bucketName string, keys []string) ([]string, error) {
	bucket, err := c.getBucket(bucketName)
	if err != nil {
		return nil, err
	}
	res, err := bucket.DeleteObjects(keys)
	if err != nil {
		return nil, err
	}
	return res.DeletedObjects, err
}
