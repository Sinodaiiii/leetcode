package medium

type LRUCache struct {
	cache []int //key
	kv    map[int]int
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{make([]int, capacity), make(map[int]int)}
	for i := range lru.cache {
		lru.cache[i] = -1
	}
	return lru
}

func (this *LRUCache) Get(key int) int {
	value, exist := this.kv[key]
	if exist {
		for i := 0; i < len(this.cache); i++ {
			if this.cache[i] == key {
				if i == 0 {
					break
				}
				copy(this.cache[1:i+1], this.cache[0:i])
				this.cache[0] = key
			}
		}
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	_, exist := this.kv[key]
	if exist == false {
		lastKey := this.cache[len(this.cache)-1]
		if lastKey > 0 {
			delete(this.kv, lastKey)
		}
		this.cache[len(this.cache)-1] = key
	}
	this.kv[key] = value
	_ = this.Get(key)
	return
}
