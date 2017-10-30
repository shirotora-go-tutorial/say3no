package main

import "fmt"

func main() {
	x := 3
	fmt.Println("x = ", x)
	// 値を渡している
	x1 := add1(x)

	fmt.Println("x + 1 = ", x1)
	// 故に変化はない
	fmt.Println("x = ", x)

	// ポインタを参照させる
	x1 = add1p(&x)
	fmt.Println("x + 1 = ", x1)
	fmt.Println("x = ", x)
}

// 引数に+1する関数
func add1(a int) int {
	a = a + 1
	return a
}

// 引数に+1する関数 ただし引数はポインタ
func add1p(a *int) int {
	*a = *a + 1
	return *a
}
