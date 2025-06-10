package hot100

// 在一个等号两侧同时操作多个指针实现操作增加/删除节点时，要考虑原地操作的问题（删除添加至同一个位置），否则可能重复赋值造成问题
//
// When performing operations on multiple pointers
//on both sides of an equal sign to add/delete nodes,
//the issue of in-place operations must be considered;
//otherwise, repeated assignments may cause problems.

type Cache struct {
	key, val   int
	prev, next *Cache
}

type LRUCache struct {
	size, curr int
	head, tail *Cache
	dict       map[int]*Cache
}

func Constructor(capacity int) LRUCache {
	head, tail := &Cache{}, &Cache{}
	head.next, tail.prev = tail, head
	return LRUCache{capacity, 0, head, tail, map[int]*Cache{}}
}

func (this *LRUCache) Get(key int) int {
	if cache, exist := this.dict[key]; exist {
		if this.head.next != cache {
			this.head.next, this.head.next.prev, cache.prev, cache.prev.next, cache.next, cache.next.prev = cache, cache, this.head, cache.next, this.head.next, cache.prev
		}
		return cache.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) >= 0 {
		this.head.next.val = value
	} else {
		if this.curr == this.size {
			delete(this.dict, this.tail.prev.key)
			this.tail.prev, this.tail.prev.prev.next = this.tail.prev.prev, this.tail
		} else {
			this.curr += 1
		}
		newCache := &Cache{key, value, this.head, this.head.next}
		this.dict[key] = newCache
		this.head.next, this.head.next.prev = newCache, newCache
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
