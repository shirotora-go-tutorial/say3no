package main

import (
	"fmt"
	"regexp"
	"bufio"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	str := stdin.Text()

	// AEGIOSZ
	// 4361052
	rep := regexp.MustCompile(`A`)
	str = rep.ReplaceAllString(str, "4")
	rep = regexp.MustCompile(`E`)
	str = rep.ReplaceAllString(str, "3")
	rep = regexp.MustCompile(`G`)
	str = rep.ReplaceAllString(str, "6")
	rep = regexp.MustCompile(`I`)
	str = rep.ReplaceAllString(str, "1")
	rep = regexp.MustCompile(`O`)
	str = rep.ReplaceAllString(str, "0")
	rep = regexp.MustCompile(`S`)
	str = rep.ReplaceAllString(str, "5")
	rep = regexp.MustCompile(`Z`)
	str = rep.ReplaceAllString(str, "2")

	fmt.Println(str) // => "1(2)3(4)5(6)"

}
