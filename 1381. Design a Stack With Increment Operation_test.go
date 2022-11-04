package leetcode

import "testing"

// ["CustomStack","push","push","pop","push","push","push","increment","increment","pop","pop","pop","pop"]
// [[3],           [1],   [2],   [],   [2],   [3],   [4],   [5,100],    [2,100],    [],   [],   [],  []]
func TestCustomStack(t *testing.T) {
	customStack := New(3)
	customStack.Push(1)                // stack becomes [1]
	customStack.Push(2)                // stack becomes [1, 2]
	t.Run("get1", func(t *testing.T) { // return 2 --> Return top of the stack 2, stack becomes [1]
		if got := customStack.Pop(); got != 2 {
			t.Errorf("Get() = %v, want %v", got, 2)
		}
	})
	customStack.Push(2)                // stack becomes [1, 2]
	customStack.Push(3)                // stack becomes [1, 2, 3]
	customStack.Push(4)                // stack still [1, 2, 3], Don't add another elements as size is 4
	customStack.Increment(5, 100)      // stack becomes [101, 102, 103]
	customStack.Increment(2, 100)      // stack becomes [201, 202, 103]
	t.Run("get1", func(t *testing.T) { // return 103 --> Return top of the stack 103, stack becomes [201, 202]
		if got := customStack.Pop(); got != 103 {
			t.Errorf("Get() = %v, want %v", got, 103)
		}
	})
	t.Run("get1", func(t *testing.T) { // return 202 --> Return top of the stack 102, stack becomes [201]
		if got := customStack.Pop(); got != 202 {
			t.Errorf("Get() = %v, want %v", got, 202)
		}
	})
	t.Run("get1", func(t *testing.T) { // return 201 --> Return top of the stack 101, stack becomes []
		if got := customStack.Pop(); got != 201 {
			t.Errorf("Get() = %v, want %v", got, 201)
		}
	})
	t.Run("get1", func(t *testing.T) { // return -1 --> Stack is empty return -
		if got := customStack.Pop(); got != -1 {
			t.Errorf("Get() = %v, want %v", got, -1)
		}
	})
}
