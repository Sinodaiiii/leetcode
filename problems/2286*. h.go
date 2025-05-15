package problems

type BookMyShow struct {
	rows       int
	seats      int
	firstEmpty []int
	totalLeft  int
}

func Constructor(n int, m int) BookMyShow {
	left := make([]int, n)
	for i := range left {
		left[i] = 0
	}
	return BookMyShow{n, m, left, n * m}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	if k > this.totalLeft {
		return nil
	}
	for i := range this.firstEmpty {
		if i > maxRow {
			break
		}
		if k <= this.seats-this.firstEmpty[i] {
			this.firstEmpty[i] = k + this.firstEmpty[i]
			this.totalLeft = this.totalLeft - k
			return []int{i, this.firstEmpty[i] - k}
		}
	}
	return nil
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if k > this.totalLeft {
		return false
	}
	var left = k
	for i := range this.firstEmpty {
		if i > maxRow {
			break
		}
		if this.firstEmpty[i] == this.seats {
			continue
		}
		left = left + this.firstEmpty[i] - this.seats
		if left <= 0 {
			for j := 0; j < i; j++ {
				this.firstEmpty[j] = this.seats
			}
			this.firstEmpty[i] = this.seats + left
			this.totalLeft = this.totalLeft - k
			return true
		}
	}
	return false
}
