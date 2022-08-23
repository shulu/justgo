package gotest

// Fib 测试方法
// @param n int
// @param m string
// @return int
// @date 2022-08-12 02:51:32
// @author shulu
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
