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

func sAtoi(input string) (m, n, k int) {
	var arr_st = []string{}
	var arr_int = []int{}
	arr_st = strings.Split(input, " ")
	for _, v := range arr_st {
		temp, _ := strconv.Atoi(v)
		arr_int = append(arr_int, temp)
	}
	return arr_int[0], arr_int[1], arr_int[2]
}

func main() {
	var m, n, k int = sAtoi(nextin()) // 立候補者m,有権者n, 演説回数k
	var who = []int{}
	for i := 0; i < m+1; i++ {
		who = append(who, 0)
	}
	who[len(who)-1] = n

	for i := 0; i < k; i++ {
		speech, _ := strconv.Atoi(nextin())
		who = move(speech-1, who)
	}

	max := 0
	for i := 0; i < len(who)-1; i++ {
		if who[i] > max {
			max = who[i]
		}
	}

	for i := 0; i < len(who)-1; i++ {
		if who[i] == max {
			fmt.Println(i+1)
		}
	}


}

func move(speech int, who []int) ([]int) {
	// 減算
	count := 0
	for i := 0; i < len(who); i++ {
		if speech == i {
			continue
		} else if who[i] >= 1 {
			who[i]--
			count++
		}
	}

	// 加算
	who[speech] += count
	return who
}
