package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	int, _ := strconv.Atoi(stdin.Text())

	for i := 1; i <= int; i++ {
		fmt.Printf("*")
	}

}
