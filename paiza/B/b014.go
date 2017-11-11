package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var sci = bufio.NewScanner(os.Stdin)

func nextIn() string {
	sci.Scan()
	return sci.Text()
}

func sAtoi(input string) (m, n, k int) {
	var arrStr = []string{}
	var arrInt = []int{}
	arrStr = strings.Split(input, " ")
	for _, v := range arrStr {
		temp, _ := strconv.Atoi(v)
		arrInt = append(arrInt, temp)
	}
	return arrInt[0], arrInt[1], arrInt[2]
}

func main() {
	var x, y, z int = sAtoi(nextIn())
	var yoko = []string{}
	var sideView = []string{}

	yoko = initYoko(yoko, y)

	for i := 1; i <= z; i++ {
		for j := 1; j <= x; j++ {
			row := splitEachChar(nextIn()) // まず一文字ずつ分割.配列の長さはy
			for i, v := range row {
				yoko = updateYoko(yoko, i, v) // 長さyのスライスを返す
			}
		}
		nextIn()
		sideView = append(sideView, concatYoko(yoko))
		yoko = initYoko(yoko, y)
	}

	for i := len(sideView) - 1; i >= 0; i-- {
		fmt.Println(sideView[i])
	}
}

func splitEachChar(str string) []string {
	return strings.Split(str, "")
}

func updateYoko(yoko []string, i int, v string) []string {
	if isFilledCell(v) && !(isFilledCell(yoko[i])) { // v が# かつ yokoが.なら更新
		yoko[i] = v
	}
	return yoko
}

func isFilledCell(str string) (ret bool) {
	if str == "#" {
		ret = true
	}
	return ret
}

func initYoko(yoko []string, y int) []string {
	yoko = nil
	for i := 1; i <= y; i++ {
		yoko = append(yoko, ".")
	}
	return yoko
}

func concatYoko(yoko []string) (concatYoko string) {
	for _, v := range yoko {
		concatYoko += v
	}
	return concatYoko
}
