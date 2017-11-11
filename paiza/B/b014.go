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
	var x, y, z int = sAtoi(nextin()) // 立候補者m,有権者n, 演説回数k
	var yy = []string{}
	var yoko = [][]string{}
	row := ""
	var outputs = []string{}

	for i := 1; i <= z; i++ {
		for i := 1; i <= x; i++ {
			s := nextin()
			c := []byte(s)
			for i := 0; i < y; i++ {
				s2 := string(c[i])
				yy = append(yy, s2)
			}
			yoko = append(yoko, yy)
			yy = nil
		}
		nextin()
	}
	fmt.Println(yoko)

	for i := 1; i <= z; i++ {
		for k := 1; k < y; k++ {
			for j := 1; j < x; j++ {
				fmt.Println("a:", i*x-k+1, " j: ", j-1)
				if yoko[i*x-k+1][j-1] == "#" {
					row += "#"
				}
			}
		}
		outputs = append(outputs, row)
		row = ""
	}

	fmt.Println(outputs)
}
