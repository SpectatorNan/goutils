package synczx

import "testing"

func TestCacheFetchNilCheck(t *testing.T) {
	key := "test_key"
	cf := NewCacheFetcher()

	res, _ := cf.Get(key, func() (interface{}, error) {
		return nil, nil
	})
	if res != nil {
		t.Errorf("Expected nil, got %v", res)
	}
	if res == nil {
		t.Logf("Successfully returned nil for key: %s", key)
	}
}
