package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	const input = "11111111222222222222222222222222"
	reader := bufio.NewReader(strings.NewReader(input))
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		fmt.Println(text)
	}

}
