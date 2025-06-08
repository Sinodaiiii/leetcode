package hard

type BookMyShow struct {
	n       int
	m       int
	minTree []int
	sumTree []int
}

func Constructor(n int, m int) BookMyShow {
	return BookMyShow{n, m, make([]int, n*4), make([]int, n*4)}
}

func (this *BookMyShow) queryMin(rowi, l, r, k int) int {
	if l == r {
		if this.minTree[rowi] <= k {
			return l
		}
		return this.n
	}
	mid := (l + r) / 2
	if this.minTree[2*rowi] <= k {
		return this.queryMin(2*rowi, l, mid, k)
	} else {
		return this.queryMin(2*rowi+1, mid+1, r, k)
	}
}

func (this *BookMyShow) querySum(rowi, l, r, l2, r2 int) int {
	if l2 <= l && r2 >= r {
		return this.sumTree[l]
	}
	mid := (l + r) / 2
	var sum int
	if mid >= l2 {
		sum += this.querySum(rowi*2, l, mid, l2, r2)
	}
	if mid+1 <= r2 {
		sum += this.querySum(rowi*2+1, mid+1, r, l2, r2)
	}
	return sum
}

func (this *BookMyShow) modify(rowi, l, r, target, k int) {
	if l == r && l == target {
		this.sumTree[rowi] = k
		this.minTree[rowi] = k
		return
	}
	mid := (l + r) / 2
	if mid >= target {
		this.modify(rowi*2, l, mid, target, k)
		this.sumTree[rowi] = this.sumTree[rowi*2] + this.sumTree[rowi*2+1]
		this.minTree[rowi] = k
	} else {
		this.modify(rowi*2+1, mid, r, target, k)
	}
	this.sumTree[rowi] = this.sumTree[rowi*2] + this.sumTree[rowi*2+1]
	this.minTree[rowi] = min(this.minTree[rowi*2], this.minTree[rowi*2+1])
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	rowNum := this.queryMin(1, 0, this.n-1, this.m-k)
	if rowNum > maxRow {
		return []int{}
	}
	seatNum := this.querySum(1, 0, this.n-1, rowNum, rowNum)
	this.modify(1, 0, this.n-1, rowNum, seatNum+k)
	return []int{rowNum, seatNum}
}

func (this *BookMyShow) doScatter(rowi, l, r, k, maxRow int) {
	if l == r {
		used := this.minTree[rowi] + k
		this.modify(1, 0, this.n-1, l, used)
	}
	mid := (l + r) / 2
	leftSum := this.querySum(rowi*2, l, mid, 0, maxRow)
	if (maxRow-l+1)*this.m-leftSum >= k {
		this.doScatter(rowi*2, l, mid, k, maxRow)
	} else {
		for i := l; i <= mid; i++ {
			this.modify(1, 0, this.n-1, i, this.m)
		}
		k -= leftSum
		this.doScatter(rowi*2+1, mid+1, r, k, maxRow)
	}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if (maxRow+1)*this.m-this.querySum(1, 0, this.n-1, 0, maxRow) < k {
		return false
	}
	this.doScatter(1, 0, this.n-1, k, maxRow)
	return true
}
