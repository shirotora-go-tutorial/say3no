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

func firstInput(input string) (times int, rule []string) {
	var arrStr = []string{}
	arrStr = strings.Split(input, " ")
	times, _ = strconv.Atoi(arrStr[0])
	rule = splitEachChar(arrStr[1])
	return times, rule
}

func main() {
	times, rule := firstInput(nextIn())
	var code = splitEachChar(nextIn()) // codeを一字ずつsplitする
	var ans string
	var idx int

	for _, v := range code {
		w := v
		if w != " " {
			for i := 1; i <= times; i++ {
				// 現在の文字が暗号の並び順で何番目か
				idx = ruleIndex(rule, w)
				w = alphabetIndex(idx)
			}
		}
		ans += w
	}

	fmt.Println(ans)

}

func splitEachChar(str string) []string {
	return strings.Split(str, "")
}

func ruleIndex(code []string, s string) (idx int) {
	for i, v := range code {
		if v == s {
			idx = i
		}
	}
	return idx
}

func alphabetIndex(idx int) (w string) {
	switch idx {
	case 0:
		w = "a"
	case 1:
		w = "b"
	case 2:
		w = "c"
	case 3:
		w = "d"
	case 4:
		w = "e"
	case 5:
		w = "f"
	case 6:
		w = "g"
	case 7:
		w = "h"
	case 8:
		w = "i"
	case 9:
		w = "j"
	case 10:
		w = "k"
	case 11:
		w = "l"
	case 12:
		w = "m"
	case 13:
		w = "n"
	case 14:
		w = "o"
	case 15:
		w = "p"
	case 16:
		w = "q"
	case 17:
		w = "r"
	case 18:
		w = "s"
	case 19:
		w = "t"
	case 20:
		w = "u"
	case 21:
		w = "v"
	case 22:
		w = "w"
	case 23:
		w = "x"
	case 24:
		w = "y"
	case 25:
		w = "z"
	}

	return w
}
