package list

type MovingAverage struct {
	data    []int
	index   int
	count   int
	cap     int
	average float64
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		data:    make([]int, size),
		average: 0,
		index:   -1,
		count:   0,
		cap:     size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.index = (this.index + 1) % this.cap
	covered := this.data[this.index]
	this.data[this.index] = val
	total := this.average * float64(this.count)
	if this.count < this.cap {
		this.count += 1
	}
	this.average = (total - float64(covered) + float64(val)) / float64(this.count)
	return this.average
}
