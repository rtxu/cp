package main

import "fmt"

type Entry struct {
	Key        int
	Val        int
	Prev, Next *Entry
}

type LRUCache struct {
	hashMap  map[int]*Entry
	list     *Entry
	len, cap int
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		hashMap: make(map[int]*Entry),
		list:    &Entry{},
		cap:     capacity,
	}
	// the first entry is the least recentyly used item
	// the last  entry is the most  recentyly used item
	cache.list.Next = cache.list // head
	cache.list.Prev = cache.list // tail
	return cache
}

func (this *LRUCache) appendEntry(entry *Entry) {
	entry.Prev = this.list.Prev
	entry.Next = this.list
	entry.Prev.Next = entry
	entry.Next.Prev = entry
	this.len++
}

func (this *LRUCache) deleteEntry(entry *Entry) *Entry {
	entry.Prev.Next = entry.Next
	entry.Next.Prev = entry.Prev
	entry.Prev = nil
	entry.Next = nil
	this.len--
	return entry
}

func (this *LRUCache) Get(key int) int {
	if entry, exist := this.hashMap[key]; exist {
		this.deleteEntry(entry)
		this.appendEntry(entry)
		return entry.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	entry := &Entry{
		Key: key,
		Val: value,
	}
	if e, exist := this.hashMap[key]; exist {
		e.Val = value
		this.deleteEntry(e)
		this.appendEntry(e)
	} else {
		this.hashMap[key] = entry
		this.appendEntry(entry)
	}

	for this.len > this.cap {
		toEvictEntry := this.deleteEntry(this.list.Next)
		delete(this.hashMap, toEvictEntry.Key)
		// fmt.Printf("evict: %d\n", toEvictEntry.Key)
	}
}

func main() {

	{
		fmt.Println("Case #1")
		cache := Constructor(2)

		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Get(1) // returns 1
		fmt.Printf("expected: 1, actual: %d\n", cache.Get(1))
		cache.Put(3, 3) // evicts key 2
		fmt.Printf("expected: -1, actual: %d\n", cache.Get(2))
		cache.Put(4, 4) // evicts key 1
		fmt.Printf("expected: -1, actual: %d\n", cache.Get(1))
		fmt.Printf("expected: 3, actual: %d\n", cache.Get(3))
		fmt.Printf("expected: 4, actual: %d\n", cache.Get(4))
	}
	{
		fmt.Println("Case #2")
		cache := Constructor(2)

		cache.Put(2, 1)
		cache.Put(1, 1)
		cache.Put(2, 3)
		cache.Put(4, 1)
		fmt.Printf("expected: -1, actual: %d\n", cache.Get(1))
		fmt.Printf("expected: 3, actual: %d\n", cache.Get(2))
	}
}
