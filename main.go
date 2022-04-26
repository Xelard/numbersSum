package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var text string
	var sum uint64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	str := strings.Split(text, " ")
	if len(str) < 2 {
		fmt.Println("not digit")
		return
	}

	for _, v := range str {
		number, err := strconv.ParseUint(v, 10, 64)
		if err != nil || number == 0 {
			fmt.Println("not digit")
			return
		}
		sum += number
	}
	fmt.Println(sum)
}
