package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()

	fmt.Printf(text)

}
