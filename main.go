package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//		line := scanner.Text()
	}

	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
