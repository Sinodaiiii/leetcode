package hard

type BookMyShow struct {
	n       int
	m       int
	minTree []int
	sumTree []int64
}

func Constructor(n int, m int) BookMyShow {
	return BookMyShow{
		n:       n,
		m:       m,
		minTree: make([]int, 4*n),
		sumTree: make([]int64, 4*n),
	}
}

func (this *BookMyShow) modify(i, l, r, index, val int) {
	if l == r {
		this.minTree[i] = val
		this.sumTree[i] = int64(val)
		return
	}
	mid := (l + r) / 2
	if index <= mid {
		this.modify(i*2, l, mid, index, val)
	} else {
		this.modify(i*2+1, mid+1, r, index, val)
	}
	this.minTree[i] = min(this.minTree[i*2], this.minTree[i*2+1])
	this.sumTree[i] = this.sumTree[i*2] + this.sumTree[i*2+1]
}

func (this *BookMyShow) queryMinRow(i, l, r, val int) int {
	if l == r {
		if this.minTree[i] > val {
			return this.n
		}
		return l
	}
	mid := (l + r) / 2
	if this.minTree[i*2] <= val {
		return this.queryMinRow(i*2, l, mid, val)
	} else {
		return this.queryMinRow(i*2+1, mid+1, r, val)
	}
}

func (this *BookMyShow) querySum(i, l, r, l2, r2 int) int64 {
	if l2 <= l && r <= r2 {
		return this.sumTree[i]
	}
	mid := (l + r) / 2
	var sum int64
	if mid >= l2 {
		sum += this.querySum(i*2, l, mid, l2, r2)
	}
	if mid+1 <= r2 {
		sum += this.querySum(i*2+1, mid+1, r, l2, r2)
	}
	return sum
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	i := this.queryMinRow(1, 0, this.n-1, this.m-k)
	if i > maxRow {
		return []int{}
	}
	used := this.querySum(1, 0, this.n-1, i, i)
	this.modify(1, 0, this.n-1, i, int(used)+k)
	return []int{i, int(used)}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	usedTotal := this.querySum(1, 0, this.n-1, 0, maxRow)
	if int64(maxRow+1)*int64(this.m)-usedTotal < int64(k) {
		return false
	}
	i := this.queryMinRow(1, 0, this.n-1, this.m-1)
	for {
		used := this.querySum(1, 0, this.n-1, i, i)
		if int64(this.m)-used >= int64(k) {
			this.modify(1, 0, this.n-1, i, int(used)+k)
			break
		}
		k -= this.m - int(used)
		this.modify(1, 0, this.n-1, i, this.m)
		i++
	}
	return true
}
