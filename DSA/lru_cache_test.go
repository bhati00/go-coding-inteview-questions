package languagebasics

import (
	"testing"
)

func TestLruCache(t *testing.T) {
	lru := NewLruCache(3)

	lru.Put(1, 10)
	lru.Put(2, 20)
	lru.Put(3, 30)

	if val := lru.Get(1); val != 10 {
		t.Errorf("Expected 10, got %d", val)
	}

	lru.Put(4, 40)

	if val := lru.Get(2); val != -1 {
		t.Errorf("Expected -1, got %d", val)
	}
}
