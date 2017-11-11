package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var sci = bufio.NewScanner(os.Stdin)

func nextin() string {
	sci.Scan()
	return sci.Text()
}

func sAtoi(input string) (p, n int) {
	var arr_st = []string{}
	var arr_int = []int{}
	arr_st = strings.Split(input, " ")
	for _, v := range arr_st {
		temp, _ := strconv.Atoi(v)
		arr_int = append(arr_int, temp)
	}
	return arr_int[0], arr_int[1]
}

func main() {
	var p, n int = sAtoi(nextin())
	temp := n % (2 * p)
	if ( 1 <= temp ) && (temp <= p ) {
		d := 2*p - right(temp)
		fmt.Println(n + d)
	} else { //ひだりかわ
		d := left(temp, p)
		fmt.Println(n - (2*p - d))
	}
}

func right(a int) (b int) {
	b = 1
	for i := 1; i < a; i++ {
		if a == 1 {
			break
		}
		b = b + 2;
	}
	return b
}

func left(a, p int) (b int) {
	b = 1
	c := 2 * p
	if a != 0 {
		for i := a; i < c; i++ { // 4,5,0
			b = b + 2;
		}
	}
	return b
}
