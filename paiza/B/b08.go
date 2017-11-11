package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var sci = bufio.NewScanner(os.Stdin)

func nextstdin() string {
	sci.Scan()
	return sci.Text()
}

func splitandAtoi(input string) []int {
	var arr_st = []string{}
	var arr_int = []int{}
	arr_st = strings.Split(input, " ")
	for _, v := range arr_st {
		temp, _ := strconv.Atoi(v)
		arr_int = append(arr_int, temp)
	}
	return arr_int
}

func main() {

	var total = []int{}
	var nm = splitandAtoi(nextstdin()) // nm[0]がn
	var sum, benefit int

	for i := 0; i < nm[1]; i++ {
		temp := splitandAtoi(nextstdin())
		for _, k := range temp {
			sum += k
		}
		total = append(total, sum)
		sum = 0
	}

	for _, i := range total {
		if i >= 1 {
			benefit += i
		}
	}

	fmt.Println(benefit)

}
