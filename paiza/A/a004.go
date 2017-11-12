package main

// 反省。
// 1. 入力値と配列のギャップで毎回無駄な時間を使っている。
// 2. データ構造の把握が遅い
// 3. 命名センス
// 4. テストで描画する時に、実際のモデルにそった形でprintすること

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

var sci = bufio.NewScanner(os.Stdin)

func nextIn() string {
	sci.Scan()
	return sci.Text()
}

type vLine struct {
	branchs []node
}

type node struct {
	toV int
	toH int
}

func sAtoi(input string) (l, n, m int) {
	var arrStr []string
	var arrInt []int
	arrStr = strings.Split(input, " ")
	for _, v := range arrStr {
		temp, _ := strconv.Atoi(v)
		arrInt = append(arrInt, temp)
	}
	return arrInt[0], arrInt[1], arrInt[2]
}

func main() {
	var amida []vLine
	var l, n, m int = sAtoi(nextIn())
	// vLineのスライスを縦線の長さl,縦線の本数nに合わせて初期化する
	amida = initAmida(amida, l, n)

	for i := 1; i <= m; i++ {
		vIdx, height, nextHeight := sAtoi(nextIn())
		drawBranch(amida, vIdx-1, height-1, nextHeight-1) // 二行目以降の入力引数を使って、無向グラフを作成する
	}

	/*
	fmt.Println("  :     0     1      2     3      4      5     6")
	for i := range amida {
		fmt.Println(i, ": ", amida[i])
	}
	*/

	// 阿弥陀開始
	for i := range amida {
		if traceAmida(amida, i) == 0 {
			fmt.Println(i + 1)
			break
		}
	}

}

// 0以外にぶつかるまでカウントを薦めていき、ぶつかったらそこに移動
func traceAmida(amida []vLine, nowV int) (goal int) {
	l := len(amida[0].branchs)
	for i := 0; i < l; i++ {
		if amida[nowV].branchs[i].toV != 9999 {
			nowV, i = amida[nowV].branchs[i].toV, amida[nowV].branchs[i].toH
		}
	}
	return nowV
}

func drawBranch(amida []vLine, vIdx, height, nextHeight int) ([]vLine) {
	amida[vIdx].branchs[height].toV = vIdx + 1
	amida[vIdx].branchs[height].toH = nextHeight

	amida[vIdx+1].branchs[nextHeight].toV = vIdx
	amida[vIdx+1].branchs[nextHeight].toH = height

	return amida
}

// TODO: vLineの構造に気をつけながらsortを使えば初期化はいらなくなるのではないか？
func initAmida(vLines []vLine, l, n int) ([]vLine) {
	var inited []node
	for i := 1; i <= n; i++ {
		for i := 1; i <= l; i++ {
			// ノード{toV: , toH}のスライスを作る
			inited = append(inited, node{9999, 0})
		}
		vLines = append(vLines, vLine{inited})
		inited = nil
	}
	return vLines
}
