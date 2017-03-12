package leetcode

import "testing"

// Test146LRUCache tests LRUCache Get() and Set()
func Test146LRUCache(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	capacity := 2
	cache := Constructor(capacity)

	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.Get(1) != 1 {
		t.Logf("1 Expected %v, but got %v", 1, cache.Get(1))
		t.Fail()
	}
	cache.Put(3, 3) // evicts key 2
	if cache.Get(2) != -1 {
		t.Logf("2 Expected %v, but got %v", -1, cache.Get(2))
		t.Fail()
	}
	cache.Put(4, 4) // evicts key 1
	if cache.Get(1) != -1 {
		t.Logf("3 Expected %v, but got %v", -1, cache.Get(1))
		t.Fail()
	}
	if cache.Get(3) != 3 {
		t.Logf("4 Expected %v, but got %v", 3, cache.Get(3))
		t.Fail()
	}
	if cache.Get(4) != 4 {
		t.Logf("5 Expected %v, but got %v", 4, cache.Get(4))
		t.Fail()
	}
}

// Test146LRUCache2 tests LRUCache Get() and Set()
func Test146LRUCache2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	capacity := 2
	cache := Constructor(capacity)

	cache.Put(2, 1)
	cache.Put(3, 2)
	if cache.Get(3) != 2 {
		t.Logf("1 Expected %v, but got %v", 2, cache.Get(3))
		t.Fail()
	}
	if cache.Get(2) != 1 {
		t.Logf("2 Expected %v, but got %v", 1, cache.Get(2))
		t.Fail()
	}
	cache.Put(4, 3) // evicts key 3
	if cache.Get(2) != 1 {
		t.Logf("3 Expected %v, but got %v", 1, cache.Get(2))
		t.Fail()
	}
	if cache.Get(3) != -1 {
		t.Logf("4 Expected %v, but got %v", -1, cache.Get(3))
		t.Fail()
	}
	if cache.Get(4) != 3 {
		t.Logf("5 Expected %v, but got %v", 3, cache.Get(4))
		t.Fail()
	}
}
