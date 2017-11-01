// struct の匿名フィールド

//実はGoは型だけの定義もサポートしています。
// これはフィールド名を書かない方法ではなく、匿名フィールドです。
// 組み込みフィールドとも呼ばれます。
// ????????  どゆこと？？？？

// 無限カプセル化できるぞい、しかも初期化せずに型だけ持たせておくことができるんだぞい…
// これはメチャクールだぞい。
// ➔フィールドの継承ができる。これ、関数も持たせられるのかな…？

package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
	phone string
}

type Student struct {
	Human // 匿名フィールド。デフォではHumanの全fieldを持つことになる。
	speciality string
	Skills
	phone string
}

func main() {
	// 学生をひとり初期化する
	taro := Student{Human{"Suzuki Taro", 23, 50,"Humanのでんわばんごう"},
		"Sparse Coding",
		Skills{"hoge", "fuga"},
		"Studentのやつ"}

	taro.Skills = append(taro.Skills, "arigataya")

	for _, v := range taro.Skills {
		fmt.Println("His Skills: ", v)
	}

	fmt.Println("His age: ", taro.age)
	fmt.Println("His weight: ", taro.weight)
	fmt.Println("His speciality: ", taro.speciality)

	taro.Skills = []string{"C++", "Swift", "TDD"}
	for _, v := range taro.Skills {
		fmt.Println("His Skills: ", v)
	}

	// 同一名のフィールドがあったらどうすんの？
	// ➔外側にあるやつが優先される
	fmt.Println("His phone : ", taro.phone)
	fmt.Println("His human phone : ", taro.Human.phone)

}






