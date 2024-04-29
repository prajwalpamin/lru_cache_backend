package dao

import "time"

type LRUCache struct {
	NodeMap map[int]Node
	Cap     int
}

type Node struct {
	value    int
	lastUsed time.Time
}

func NewLRUCache(c int) *LRUCache {
	return &LRUCache{
		Cap:     c,
		NodeMap: make(map[int]Node, c),
	}
}

func (lru *LRUCache) Put(key, value int) {
	mSize := len(lru.NodeMap)
	var leastRecentUsedKey int
	leastRecentUsedValue := time.Now()

	// remove if the max size reached
	if mSize >= lru.Cap {
		// remove most recent one
		for key, v := range lru.NodeMap {
			if v.lastUsed.Before(leastRecentUsedValue) {
				leastRecentUsedKey = key
				leastRecentUsedValue = v.lastUsed
			}
		}

		delete(lru.NodeMap, leastRecentUsedKey)
	}

	lru.NodeMap[key] = Node{
		value:    value,
		lastUsed: time.Now(),
	}
}

// Get ...
func (lru *LRUCache) Get(key int) int {
	val, prs := lru.NodeMap[key]

	if prs {
		return val.value
	}
	return -1
}
