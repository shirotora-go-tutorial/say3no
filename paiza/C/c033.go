package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var sci = bufio.NewScanner(os.Stdin)

func nextIn() string {
	sci.Scan()
	return sci.Text()
}

type ballCount struct {
	ball   int
	strike int
}

func main() {
	// 標準入力の処理
	var bc = ballCount{0, 0}
	var ballTypes []string
	pitches, _ := strconv.Atoi(nextIn())

	for i := 1; i <= pitches; i++ {
		ballTypes = append(ballTypes, nextIn())
	}

	for i := 0; i < pitches; i++ {
		if isStrike(ballTypes[i]) {
			bc.strike++
		} else {
			bc.ball++
		}
		shoutJudgement(bc, ballTypes[i])
	}
}

func shoutJudgement(bc ballCount, ballType string) {
	if bc.ball == 4 {
		fmt.Println("fourball!")
	} else if bc.strike == 3 {
		fmt.Println("out!")
	} else {
		fmt.Println(ballType + "!")
	}
}

func isStrike(bType string) bool {
	var j bool
	if bType == "strike" {
		j = true
	}
	return j
}
