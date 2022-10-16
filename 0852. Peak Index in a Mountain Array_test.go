package leetcode

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"1", []int{0, 1, 0}, 1},
		{"2", []int{0, 2, 1, 0}, 1},
		{"3", []int{0, 10, 5, 2}, 1},
		{"4", []int{0, 1, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args); got != tt.want {
				t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
