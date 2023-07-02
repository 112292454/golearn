package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	var tempMap = make(map[string]int)
	var flag = false

	input := bufio.NewScanner(f)
	for input.Scan() {
		var s = input.Text()
		if tempMap[s] >= 1 {
			flag = true
		}
		tempMap[s]++
	}
	if flag {
		fmt.Printf("在文件%s中发现重复行\n", f.Name())
	}
	for k, v := range tempMap {
		counts[k] += v
	}
}
