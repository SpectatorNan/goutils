package cachex

import (
	"golang.org/x/sync/singleflight"
	"sync"
)

type CacheFetcher struct {
	cacheMap sync.Map
	sfGroup  singleflight.Group
}

func NewCacheFetcher() *CacheFetcher {
	return &CacheFetcher{
		cacheMap: sync.Map{},
		sfGroup:  singleflight.Group{},
	}
}
func (cf *CacheFetcher) Get(key string, fetchFunc func() (interface{}, error)) (interface{}, error) {
	if v, ok := cf.cacheMap.Load(key); ok {
		return v, nil
	}

	v, err, _ := cf.sfGroup.Do(key, func() (interface{}, error) {
		if v, ok := cf.cacheMap.Load(key); ok {
			return v, nil
		}
		v, err := fetchFunc()
		if err != nil {
			return nil, err
		}
		cf.cacheMap.Store(key, v)
		return v, nil
	})
	return v, err
}

func (cf *CacheFetcher) Delete(key string) {
	cf.cacheMap.Delete(key)
}

func (cf *CacheFetcher) Flush() {
	cf.cacheMap = sync.Map{}
}
