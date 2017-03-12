package leetcode

// LRUCache thing
type LRUCache struct {
	m        map[int]*cacheEntry
	first    *cacheEntry
	last     *cacheEntry
	size     int
	capacity int
}

type cacheEntry struct {
	prev  *cacheEntry
	next  *cacheEntry
	value int
	key   int
}

// Constructor to create LRUCache
func Constructor(capacity int) LRUCache {
	println("************")
	c := LRUCache{map[int]*cacheEntry{}, nil, nil, 0, capacity}
	return c
}

// Get retrieve value by key of LRUCache
func (cache *LRUCache) Get(key int) int {
	entry, ok := cache.m[key]
	if ok {
		println("found", key)
		if entry.prev != nil {
			entry.prev.next = entry.next
			println("prev to next")
		} else {
			cache.first = entry.next
			println("first is next")
			if cache.first == nil {
				println("first is nil")
			}
		}
		if entry.next != nil {
			entry.next.prev = entry.prev
			println("next to prev")
		} else {
			cache.last = entry.prev // could check at the start, if we are last nothing to be done
		}
		entry.next = nil
		entry.prev = cache.last
		if cache.last != nil {
			cache.last.next = entry
			println("last to me")
		}
		cache.last = entry

		if cache.first == nil {
			cache.first = entry
			println("first to me") // may be a better place for this
		}
		return entry.value
	}
	return -1
}

// Put value of LRUCache
func (cache *LRUCache) Put(key int, value int) {
	existing := cache.Get(key)
	if existing < 0 {

		println("*** insert", key)
		entry := cacheEntry{cache.last, nil, value, key}
		if cache.last != nil {
			cache.last.next = &entry
			println("last to me 2")
		}
		cache.last = &entry
		cache.m[key] = &entry
		if cache.first == nil {
			cache.first = &entry
		}

		cache.size++

		println("size", cache.size)

		if cache.size > cache.capacity {
			toDelete := cache.first
			cache.first = cache.first.next
			if cache.first != nil {
				cache.first.prev = nil
			}
			delete(cache.m, toDelete.key)
			println("evict", toDelete.key)
			cache.size--
		}
	} else {
		entry := cache.m[key] // could optimise away as it is called in Get() above
		entry.value = value
	}
}
