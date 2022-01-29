package leetcode

import "testing"

func clabFloor(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var f [10 + 1]int
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]

}

func TestClabFloor(t *testing.T) {
	t.Log(clabFloor(5))
}
