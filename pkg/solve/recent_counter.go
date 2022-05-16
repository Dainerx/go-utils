package solve

const threeSec = 3000

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	q := make([]int, 0)
	return RecentCounter{q}
}

func (this *RecentCounter) push(t int) {
	this.q = append(this.q, t)
}

func (this *RecentCounter) pop() int {
	r := this.q[0]
	this.q = this.q[1:]
	return r
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.q) > 0 && t-this.q[0] > threeSec {
		this.pop()
	}
	this.push(t)
	return len(this.q)
}
