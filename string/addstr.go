package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "hello"
	i1, err := strconv.ParseInt(str1, 6, 12)

	fmt.Println(i1)

	str2 := "hihello"
	i2, err := strconv.ParseInt(str2, 5, 20)
	if err == nil {
		fmt.Println(i2)
	}
}
