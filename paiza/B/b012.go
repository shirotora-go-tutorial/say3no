package main

import (
	"strconv"
	"bufio"
	"os"
	"fmt"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := nextLine()
	digits := []string{}
	num, _ := strconv.Atoi(s)
	sums := []int{}
	for i := 0; i < num; i++ {
		digits = append(digits, nextLine())
	}

	for i := 0; i < num; i++ {
		sums = append(sums, checkDigit(digits[i]))
	}

	for i := 0; i < num; i++ {
		if ( 10 - (sums[i] % 10) ) == 10 {
			fmt.Println(0)
		} else {
			fmt.Println(10 - (sums[i] % 10))
		}
	}

}

func checkDigit(digits string) int {
	evenSum := 0
	oddSum := 0
	for pos, v := range digits {
		if pos == 16 {
			break
		}

		if pos%2 == 0 {
			ev, _ := strconv.Atoi(string([]rune{v}))
			switch ev {
			case 5:
				ev = 1
			case 6:
				ev = 3
			case 7:
				ev = 5
			case 8:
				ev = 7
			case 9:
				ev = 9
			default:
				ev = 2 * ev
			}
			evenSum += ev
		} else {
			od, _ := strconv.Atoi(string([]rune{v}))
			oddSum += od
		}
	}
	return oddSum + evenSum
}
