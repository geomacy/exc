package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func process(file string) (map[string]int, error) {
	counts := make(map[string]int, 0)
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	entries := strings.Fields(string(contents))
	for _, e := range entries {
		if e != "" {
			if _, ok := counts[e]; !ok {
				counts[e] = 0
			}
			counts[e] = counts[e] + 1
		}
	}
	return counts, nil
}

func main() {
	counts, err := process(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
	fmt.Printf("%v\n", counts)
}
