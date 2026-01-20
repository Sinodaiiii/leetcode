package hard

type Cache struct {
	key       int
	value     int
	heapIndex int
	times     int
	version   int
}

type LFUCache struct {
	dict         map[int]*Cache
	heap         []*Cache
	capacity     int
	totalVersion int
}

func Constructor260120(capacity int) LFUCache {
	return LFUCache{map[int]*Cache{}, make([]*Cache, 0, capacity), capacity, 0}
}

func (this *LFUCache) adjustHeap(i int) {
	tmp := i
	if i*2+1 < len(this.heap) && (this.heap[tmp].times > this.heap[i*2+1].times || (this.heap[tmp].times == this.heap[i*2+1].times && this.heap[tmp].version > this.heap[i*2+1].version)) {
		tmp = i*2 + 1
	}
	if i*2+2 < len(this.heap) && (this.heap[tmp].times > this.heap[i*2+2].times || (this.heap[tmp].times == this.heap[i*2+2].times && this.heap[tmp].version > this.heap[i*2+2].version)) {
		tmp = i*2 + 2
	}
	if tmp != i {
		this.heap[tmp], this.heap[i] = this.heap[i], this.heap[tmp]
		this.heap[tmp].heapIndex = tmp
		this.heap[i].heapIndex = i
		this.adjustHeap(tmp)
	}
}

func (this *LFUCache) heapUp(i int) {
	for i > 0 {
		if this.heap[i].times < this.heap[(i-1)/2].times || (this.heap[i].times == this.heap[(i-1)/2].times && this.heap[i].version < this.heap[(i-1)/2].version) {
			this.heap[i], this.heap[(i-1)/2] = this.heap[(i-1)/2], this.heap[i]
			this.heap[(i-1)/2].heapIndex = (i - 1) / 2
			this.heap[i].heapIndex = i
			i = (i - 1) / 2
		} else {
			break
		}
	}
	return
}

func (this *LFUCache) Get(key int) int {
	if v, exist := this.dict[key]; exist {
		this.totalVersion += 1
		v.version = this.totalVersion
		v.times += 1
		this.adjustHeap(v.heapIndex)
		return v.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	this.totalVersion += 1
	if v, exist := this.dict[key]; exist {
		v.version = this.totalVersion
		v.times += 1
		v.value = value
		this.adjustHeap(v.heapIndex)
	} else if len(this.heap) == this.capacity {
		oldKey := this.heap[0].key
		this.heap[0] = &Cache{key, value, 0, 1, this.totalVersion}
		delete(this.dict, oldKey)
		this.dict[key] = this.heap[0]
		this.adjustHeap(0)
	} else {
		this.dict[key] = &Cache{key, value, len(this.heap), 1, this.totalVersion}
		this.heap = append(this.heap, this.dict[key])
		this.heapUp(len(this.heap) - 1)
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
