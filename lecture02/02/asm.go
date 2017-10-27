package main

import "fmt"

// var arr [n]type
var arr [10]int // intは0で初期化される？
//var brr [5]int // intは0で初期化される？
// 配列長は変えることができない。また、配列長が異なると、それぞれ違う型として扱わなければならない。
// つまり arr = brr みたいな代入は出来ないし、[5:10]の部分は据え置きとかパディングとか、そういうこともしてくれない。（もしかしたらそういう関数は用意されているのかもしれないが）
var fslcie []int //slice

func main() {
	//ふつーの配列
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0]) // データを取得して、42を返します。
	fmt.Printf("The last element is %d\n", arr[9])  //値が代入されていない最後の要素を返します。デフォルトでは0が返ります。

	/*
	unpoko := [2]int{1, 2}            // 短縮宣言でも配列の場合は型宣言必須なん？
	ahoaho := [...]int{3, 4, 5, 6, 7} // サイズを要素数から判別してくれる
	bakabaka := [10]int{8, 9, 10, 11} // 配列長>要素数のときは、頭から詰めていれてくれる。
	*/
	// この場合、{8,9,10,11,0,0,0,0,0,0}となる

	/*
	// 配列のネストも勿論可能…！
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}} // doubleArrayと等価
	*/

	// slice
	// 可変長の配列とも言うべきGo独特のデータ構造。動的な配列**っぽい**動作をする。
	// 実際は*常に低レイヤのarrayを指している*。単なる配列の参照型なのである。
	//	slice := []byte{'a', 'b', 'o', 'i', 'c'}

	// 要するに、配列長を指定せずに配列を作成すればそいつはsliceになるっぽい

	/* 例 */
	var ar = [10]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yz!"}

	// byteを含む２つのsliceを宣言します
	var a, b []string

	a = ar[2:5]
	b = ar[3:5]
	fmt.Println(a)
	fmt.Println(b)

	ar[4] = "あいうえお"

	fmt.Println(a)
	fmt.Println(a.len)
	fmt.Println(a.)
	fmt.Println(b)

	/* 重要なのは、sliceは参照型だ、ということ */

}
