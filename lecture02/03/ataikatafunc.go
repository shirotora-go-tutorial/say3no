/*
 Goでは関数も変数の一種です。 typeを通して定義します。<T>みたいな。
 これは、全て同じ引数と同じ戻り値を持つ一つの型です。
 type typeName func(input1 inputType1, input2 inputType2 [, ...]) (result1 resultType1 [, ...])
 この辺もjsっぽいよね

 関数を方として扱うことにメリットは有るのでしょうか？ では、この型の関数を値として渡してみましょう。以下の例をご覧ください。
*/

package main

import "fmt"

type testInt func(int) bool // 関数の型を宣言します
/*
func(int) bool 型の関数を扱う型。
ちなみに構造体の定義は多分typeでやるんだと思う
 */

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// ここでは宣言する関数の型を引数のひとつとみなします。
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 関数の値渡し
	fmt.Println("Odd elements of slice are : ", odd)
	even := filter(slice, isEven) // 関数の値渡し
	fmt.Println("Odd elements of slice are : ", even)

}
