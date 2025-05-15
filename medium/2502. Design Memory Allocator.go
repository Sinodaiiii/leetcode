package medium

import "sort"

type Allocator struct {
	free      [][2]int
	allocated map[int][][2]int
}

func Constructor(n int) Allocator {
	return Allocator{[][2]int{{0, n - 1}}, make(map[int][][2]int)}
}

func (this *Allocator) Allocate(size int, mID int) int {
	if this.free[0][0] == -1 {
		return -1
	}
	ans := 0
	for curr := 0; curr < len(this.free); curr++ {
		if this.free[curr][1]-this.free[curr][0]+1 >= size {
			_, exist := this.allocated[mID]
			if !exist {
				this.allocated[mID] = make([][2]int, 0)
			}
			this.allocated[mID] = append(this.allocated[mID], [2]int{this.free[curr][0], this.free[curr][0] + size - 1})
			sort.Slice(this.allocated[mID], func(a, b int) bool {
				return this.allocated[mID][a][0] < this.allocated[mID][b][0]
			})
			ans = this.free[curr][0]
			this.free[curr][0] += size
			if this.free[curr][0] >= this.free[curr][1] {
				if len(this.free) == 1 {
					this.free[0][0] = -1
				} else {
					this.free = append(this.free[:curr], this.free[curr+1:]...)
				}
			}
			return ans
		}
	}
	return -1
}

func (this *Allocator) FreeMemory(mID int) int {
	target, exist := this.allocated[mID]
	if !exist {
		return 0
	}
	total := 0
	for _, block := range target {
		total += block[1] - block[0] + 1
		for i := 0; i < len(this.free); i++ {
			if this.free[0][0] == -1 {
				this.free[0] = block
			} else {
				if block[0] <= this.free[i][0] {
					this.free = append(this.free, [2]int{})
					copy(this.free[i+1:], this.free[i:])
					this.free[i] = block
					for j := 0; j < len(this.free)-1; j++ {
						if this.free[j][1]+1 == this.free[j+1][0] {
							this.free[j+1][0] = this.free[j][0]
							this.free = append(this.free[:j], this.free[j+1:]...)
						}
					}
					break
				}
				//if i >= 1 {
				//	if block[0] == this.free[i-1][1] + 1 {
				//		this.free[i-1][1] = block[1]
				//	}
				//}
				//if block[1] + 1 == this.free[i][1] {
				//	this.free[]
				//}
				//if block[0] < this.free[i][0] {
				//	if block[1]+1 == this.free[0][0] {
				//		this.free[0][0] = block[0]
				//	} else {
				//		this.free = append([][2]int{block}, this.free...)
				//	}
				//	break
				//}
			}
		}
	}
	delete(this.allocated, mID)
	return total
}
