// 实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
package main

func fibonacci() func() int {
	x, y, z := 0, 1, 0
	return func() int {
		z, x, y = x, y, x+y
		return z
	}
}

// func main() {
// 	f := fibonacci()

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }
