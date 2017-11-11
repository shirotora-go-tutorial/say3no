package main

import "fmt"
import "regexp"

func main() {
	str := "rssssss"
	rep := regexp.MustCompile(`s$`)
	str = rep.ReplaceAllString(str, "A")

	fmt.Println(str) // => "Copyleft 2015 by ASHITANI Tatsuji."
}
