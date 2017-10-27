package main

import "fmt"

var frenchHello string      // 文字列変数の宣言の一般的な方法
var emptyString string = "" //空文字列で初期化

func main() {
	// 変数
	// var variableName type
	// var vname1, vname2, vname3 type
	// 変数宣言、変数名、型、値の順。常に変わらない？
	// var variableName type = value
	// var vname1, vname2, vname3 type = v1,v2,v3
	// 直接型の宣言を無視することができるので、上のコードはこのようにも書けます：
	// var vname1, vname2, vname3 = v1, v2, v3
	//これでもまだ面倒ですか？ええ、私もそう思います。更に簡単にしてみましょう。
	// hoge := v1　。これを短縮宣言と呼ぶ。が、関数内でしか使えない。まあそりゃそうだな。
	// グローバル変数はvar方式で型等を明示する

	//不要な変数など存在しない。それがGO!使っていない変数があるとコンパイルエラーになる。

	//	vname1, vname2, vname3 := _, _, _
	fmt.Printf("hogehoge")

	// 定数
	//	const constantName = value
	//	const constantName type = value

	// 型
	// ビルトイン基本型

	/* boolean
	var isActive bool = false
	var enabled, disabled = true, false
	func test(){
		var available bool //通常の宣言
		valid := false   // 短縮宣言
		available = true //代入操作
	}
	*/

	// 数値型
	// int, uint。ちなuintは非負のみ。intは負も扱う。
	// それぞれは別の型であり、int32+int8みたいな演算はできない。
	// 浮動小数点は`float32`と`float64`の二つのみ。ただの`float`は存在しない。

	// 複素数型！
	// デフォルトは`complex128`型。64bit実数と64bit虚数からなる。

	var z complex64 = 5 + 5i
	// output: (5+5i)
	fmt.Printf("Value is: %v \n", z)

	// 文字列は""か``。シングルクォートじゃなくてバッククォートだよ。

	test()

	// Goの文字列は変更することができない（マジ？）
	/*
	これはコンパイルエラー
		var s string = "hello"
		s[0] = 'c'
	 */

	// どーしても変更したくなったらbyte型にキャストしてからやる。
	// その後、String型に戻す。
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	// +演算子で文字列を連結

	s = "hogehoge, "
	m := " unpoko"
	a := s + m
	fmt.Printf("%s\n", a)

	// スライスもできる
	s = "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)

	// 複数行の文字列を宣言したい。ただしRaw文字列となる点が注意
	m = `hellow
worlda
    sdlkfjlksdjfa
                dsafkjlsdfa`
	fmt.Printf("%s\n", m)

}

func test() {
	no, yes, maybe := "no", "yes", "maybe" // 短縮宣言、同時に複数の変数を宣言
	japaneseHello := "Konichiwa"           // 同上
	frenchHello = "Bonjour"                //通常の代入
	fmt.Printf("%s,%s,%s,%s,%s", no, yes, maybe, japaneseHello, frenchHello)
}
