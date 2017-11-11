package main

import (
	"bufio"
	"os"
	"strconv"
	"regexp"
	"fmt"
)

var sci = bufio.NewScanner(os.Stdin)

func nextstdin() string {
	sci.Scan()
	return sci.Text()
}

func main() {

	n, _ := strconv.Atoi(nextstdin())
	var words = []string{}

	for i := 1; i <= n; i++ {
		words = append(words, nextstdin())
	}

	pl := ""
	for _, i := range words {
		// s, sh ch, o, x ➔ es(追加)
		rep := regexp.MustCompile(`s$|sh$|ch$|o$|x$`)
		if (rep.MatchString(i) == true) {
			fmt.Println(i + "es")
			continue
		}
		// f, fe ➔ vesに置換
		rep = regexp.MustCompile(`f$|fe$`)
		if (rep.MatchString(i) == true) {
			pl = rep.ReplaceAllString(i, "ves")
			fmt.Println(pl)
			continue
		}

		// [aiueo意外]y ➔ [aiueo意外]ies
		rep = regexp.MustCompile(`[^aiueo]y$`)
		if (rep.MatchString(i) == true) {
			rep = regexp.MustCompile(`y$`)
			pl = rep.ReplaceAllString(i, "ies")
			fmt.Println(pl)
			continue
		}

		// 上記意外➔s(追加)
		fmt.Println(i + "s")
		continue
	}
}

