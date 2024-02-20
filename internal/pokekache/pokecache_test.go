package pokekache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
    cache := NewCache(time.Millisecond)
    if cache.cache == nil {
        t.Error("Cache was not created")
    }
}

func TestSetGetCache(t *testing.T) {
    cache := NewCache(time.Millisecond)

    cases := []struct {
        key string
        val []byte
    }{
        {"key1", []byte("val1")},
        {"key2", []byte("val2")},
        {"key3", []byte("val3")},
    }
    for _, c := range cases {
        cache.Set(c.key, []byte(c.val))
        actual, ok := cache.Get(c.key)
        if !ok {
            t.Errorf("%s key not found in cache", c.key)
        }
        if string(actual) != string(c.val) {
            t.Errorf("Expected %s, got %s", c.val, string(actual))
        }
    }
}

func TestReapLoopCache(t *testing.T) {
    interval := time.Millisecond * 10
    cache := NewCache(interval)

    cases := []struct {
        key string
        val []byte
    }{
        {"key1", []byte("val1")},
        {"key2", []byte("val2")},
        {"key3", []byte("val3")},
    }
    for _, c := range cases {
        cache.Set(c.key, []byte(c.val))
    }

    time.Sleep(interval * 2)

    for _, c := range cases {
        _, ok := cache.Get(c.key)
        if ok {
            t.Errorf("%s key not deleted from cache", c.key)
        }
    }
}
