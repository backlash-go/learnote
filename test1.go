package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()] = counts[input.Text()] + 1 //counts[input.Text()] ++
	}

	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {

			fmt.Printf("------- %d\t%s\n", n, line)

	}
}
