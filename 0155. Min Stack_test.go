package leetcode

import "testing"

// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
func TestMinStack1(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Run("getMin1", func(t *testing.T) { // return -3
		if got := minStack.GetMin(); got != -3 {
			t.Errorf("Get() = %v, want %v", got, -3)
		}
	})
	minStack.Pop()
	t.Run("top", func(t *testing.T) { // return 0
		if got := minStack.Top(); got != 0 {
			t.Errorf("Get() = %v, want %v", got, 0)
		}
	})
	t.Run("getMin2", func(t *testing.T) { // return -2
		if got := minStack.GetMin(); got != -2 {
			t.Errorf("Get() = %v, want %v", got, -2)
		}
	})
}

// ["MinStack","push","push","push","getMin","pop","pop","top","getMin"]
// [[],[5],[3],[4],[],[],[],[],[]]
func TestMinStack2(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(5)
	minStack.Push(3)
	minStack.Push(4)
	t.Run("getMin1", func(t *testing.T) { // return 3
		if got := minStack.GetMin(); got != 3 {
			t.Errorf("Get() = %v, want %v", got, 3)
		}
	})
	minStack.Pop()
	minStack.Pop()
	t.Run("top", func(t *testing.T) { // return 5
		if got := minStack.Top(); got != 5 {
			t.Errorf("Get() = %v, want %v", got, 5)
		}
	})
	t.Run("getMin2", func(t *testing.T) { // return 5
		if got := minStack.GetMin(); got != 5 {
			t.Errorf("Get() = %v, want %v", got, 5)
		}
	})
}
