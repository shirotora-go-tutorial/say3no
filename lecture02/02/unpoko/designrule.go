package unpoko

import "fmt"

// 大文字で始まる変数はエクスポート可能。つまり他のパッケージから読み込み可能。public.
const (
	Age    = 29
)

// 小文字ではじまる変数はエクスポート不可能
const (
	mind      = "happy"
	knowledge = 100
)

func Public() int { //外から呼べる
	fmt.Printf(private())
	return Age
}

func private() string {
	return "private" + mind
}
