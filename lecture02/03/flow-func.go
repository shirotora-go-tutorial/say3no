package main

import "fmt"

func main() {
	// if に括弧は要らない。Goで括弧を使うのは明らかな引数にだけなのかな・・・?
	if x := computedValue(); x > 10 { // xのスコープはifの中だけ
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// for; Goでもっとも強力なロジックコントロール（って書いてありました)
	// c likeなfor文っぽい

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Printf("sum is equal to %v \n", sum)

	// 初期条件、終了時処理を省略することができる
	// つまり、条件だけを記述することができる（それもうwhile)やん
	hoge := 1
	for ; hoge < 1000; {
		hoge += hoge
	}
	fmt.Printf("hoge is equal to %v \n", hoge)

	// 初期条件、終了時処理を省略するとき、; を省略できる
	// つーかwhileやん
	piyo := 1
	for piyo < 1000 {
		piyo += piyo
	}
	fmt.Printf("hoge is equal to %v \n", piyo)

	// break, continue
	for unpoko := 10; unpoko > 0; unpoko-- {
		if unpoko == 5 {
			break // またはcontinue
		}
		fmt.Println(unpoko)
	}

	uouo := make(map[string]string)
	uouo["キー"] = "ばりゅー"
	uouo["key"] = "value"

	// forはrangeと組み合わせて、array,slice,map,stringのデータを読み込ませることとができる。
	// pythonでつかうrangeを取り込んで、foreachとかを一般化したのかな
	for k, v := range uouo {
		fmt.Println("map's key:", k)
		fmt.Println("map's value:", v)
	}

	// Goは複数の戻り値をサポートしているが、宣言して仕様されていない変数に対してコンパイルはエラーを出力する。
	// この時、_ を使うことで必用のない戻り値を切り捨てることができる。
	for _, v := range uouo {
		fmt.Println("map's value:", v)
	}

	//switch
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("確かなことはiはintだってこと")
	}

	// fallthroufh in switch
	// 以下のケースを全て強制実行させる
	// ただしdefaultは実行しない
	j := 6
	switch j {
	case 1:
		fmt.Println("j is equal to 1")
	case 2, 3, 4:
		fmt.Println("j is equal to 2, or 4")
	case 5, 6:
		fmt.Println(" iusasdkjfalsd")
		fallthrough
	case 10:
		fmt.Println("j is equal to 10")
	default:
		fmt.Println("確かなことはiはintだってこと")
	}

	x, y, z := 3, 4, 5
	maxXy := max(x, y)
	maxXz := max(x, z)

	fmt.Printf("max(%d, %d) = %d \n", x, y, maxXy)
	fmt.Printf("max(%d, %d) = %d \n", x, z, maxXz)
	fmt.Printf("max(%d, %d) = %d \n", y, z, max(y, z))

	a, b := 29, 30
	plus, product := SumAndProduct(a, b)
	fmt.Printf("%d + %d = %d \n", a, b, plus)
	fmt.Printf("%d * %d = %d \n", a, b, product)

	kahencho(1, 2, 3, 4, 5, 6, 7, 8, 9)
	kahencho(10,11,12,13,14)





}

/*
func funcName(input1 type1, input2 type2) (output1 type1, output2 type2){
	// ここはロジック処理のコードです。
	// 複数の値を戻り値とします。
	return value1,value2
}
*/

// a,bの中から最大値を返す
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 可変長引数
// argはintのsliceとなる
func kahencho(arg ...int) {
	for _, n := range arg {
		fmt.Printf("and the number is: %d\n", n)
	}
}

// 複数の戻り値
// A+BとA*Bを返す
// 返り値には名前をつけろ、どっちもintじゃどっちが掛け算かわからんからな
func SumAndProduct(A, B int) (Add int, Multiplied int) {
	return A + B, A * B
}

func computedValue() int {
	return 10
}

// goto文。闇
func myFunc() {
	i := 0
Here: //末尾にコロンをつけるとタグになる
	println(i)
	i++
	goto Here //Hereにジャンプする
}
