package list

import "testing"

func TestList(t *testing.T) {
	obj := MyCircularQueueConstructor(10)
	for i := 1; i < 13; i++ {
		t.Log(obj.EnQueue(i))
	}
	t.Log(obj.DeQueue())
	t.Log(obj.Front())
	t.Log(obj.Rear())
	t.Log(obj.IsEmpty())
	t.Log(obj.IsFull())
}

func TestList2(t *testing.T) {
	obj := MyCircularQueueConstructor(6)
	t.Log(obj.EnQueue(6))
	t.Log(obj.Rear())
	t.Log(obj.Rear())
	t.Log(obj.DeQueue())
	t.Log(obj.EnQueue(5))
	t.Log(obj.Rear())
	t.Log(obj.DeQueue())
	t.Log(obj.Front())
	t.Log(obj.DeQueue())
	t.Log(obj.DeQueue())
	t.Log(obj.DeQueue())
}
func TestList3(t *testing.T) {
	obj := Constructor(6)
	t.Log(obj.Next(3))
	t.Log(obj.Next(7))
	t.Log(obj.Next(9))
}
