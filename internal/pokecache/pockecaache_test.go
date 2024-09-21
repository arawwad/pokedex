package pockecache

import (
	"testing"
	"time"
)

func TestCache_AddAndGet(t *testing.T) {
	cache := NewCache(1 * time.Second)
	defer cache.Stop()

	cache.Add("v1", []byte("v1"))

	value, ok := cache.Get("v1")

	if !ok {
		t.Errorf("can not find cache item")
	}

	if string(value) != "v1" {
		t.Errorf("Expect v1 found %s", string(value))
	}

}

func TestCache_GetNonExisting(t *testing.T) {
	cache := NewCache(1 * time.Second)
	defer cache.Stop()

	_, ok := cache.Get("v1")

	if ok {
		t.Errorf("found cache value that doesn't exist")
	}

}

func TestCache_CacheExpiry(t *testing.T) {
	cache := NewCache(1 * time.Second)
	defer cache.Stop()

	cache.Add("v1", []byte("v1"))

	time.Sleep(2 * time.Second)

	_, ok := cache.Get("v1")

	if ok {
		t.Errorf("found cache value that doesn't exist")
	}
}
