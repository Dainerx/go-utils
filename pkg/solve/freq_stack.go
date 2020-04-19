package solve

type FreqStack struct {
	freq     map[int]int
	freqData map[int][]int
	maxF     int
}

func Constructor() FreqStack {
	freq := make(map[int]int)
	freqd := make(map[int][]int)
	maxF := 0
	return FreqStack{freq, freqd, maxF}
}

func (this *FreqStack) Push(x int) {
	this.freq[x]++
	f := this.freq[x]
	this.freqData[f] = append(this.freqData[f], x)
	if f >= this.maxF {
		this.maxF = f
	}
}

func (this *FreqStack) Pop() int {
	l := len(this.freqData[this.maxF])
	r := this.freqData[this.maxF][l-1]

	this.freq[r]--
	this.freqData[this.maxF] = this.freqData[this.maxF][:l-1]

	if len(this.freqData[this.maxF]) == 0 {
		this.maxF--
	}
	return r
}
