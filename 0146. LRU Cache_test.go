package leetcode

import (
	"testing"
)

// ["LRUCache","put","put","get","put","get","put","get","get","get"]
// [    [2],   [1,0],[2,2],[1],  [3,3],[2],  [4,4],[1],   [3],  [4]]
// [null,null,null,0,null,-1,null,-1,3,4]
func TestLRUCache(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 0)
	c.Put(2, 2)
	t.Run("get1", func(t *testing.T) {
		if got := c.Get(1); got != 0 {
			t.Errorf("Get() = %v, want %v", got, 0)
		}
	})
	c.Put(3, 3)
	t.Run("get2", func(t *testing.T) {
		if got := c.Get(2); got != -1 {
			t.Errorf("Get() = %v, want %v", got, -1)
		}
	})
	c.Put(4, 4)
	t.Run("get3", func(t *testing.T) {
		if got := c.Get(1); got != -1 {
			t.Errorf("Get() = %v, want %v", got, -1)
		}
	})
	t.Run("get4", func(t *testing.T) {
		if got := c.Get(3); got != 3 {
			t.Errorf("Get() = %v, want %v", got, 3)
		}
	})
	t.Run("get5", func(t *testing.T) {
		if got := c.Get(4); got != 4 {
			t.Errorf("Get() = %v, want %v", got, 3)
		}
	})
}
