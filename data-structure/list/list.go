package list

type MyCircularQueue struct {
	data  []int
	head  int
	count int
	cap   int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func MyCircularQueueConstructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, k),
		head:  0,
		count: 0,
		cap:   k,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	//检查队列是否满
	if this.IsFull() {
		return false
	}
	this.count += 1
	this.data[this.TailIndex()] = value
	return true
}
func (this *MyCircularQueue) TailIndex() int {
	return (this.head + this.count - 1) % this.cap
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.count -= 1
	this.head = (this.head + 1) % this.cap
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.TailIndex()]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.count <= 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.count == this.cap
}
