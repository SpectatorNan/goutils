package oss

import (
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func (c *OSSClient) listObjects(ctx context.Context, bucketName string, prefix string, marker string) (*oss.ListObjectsResult, error) {
	bucket, err := c.getBucket(bucketName)
	if err != nil {
		return nil, err
	}
	objects, err := bucket.ListObjects(oss.Prefix(prefix), oss.Marker(marker))
	if err != nil {
		return nil, err
	}
	return &objects, nil
}

func (c *OSSClient) FetchFileSize(ctx context.Context, buckName string, prefix string, marker string) (int64, error) {
	pref := prefix
	mark := marker
	bsize := int64(0)
	for {
		lsRes, err := c.listObjects(ctx, buckName, pref, mark)
		if err != nil {
			return 0, err
		}
		for _, obj := range lsRes.Objects {
			bsize += obj.Size
		}
		if lsRes.IsTruncated {
			mark = lsRes.NextMarker
			pref = lsRes.Prefix
		} else {
			break
		}
	}
	return bsize, nil
}
