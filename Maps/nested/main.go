package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		firstChar := []rune(name)[0]
		if _, ok := counts[firstChar]; !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}
	fmt.Println(counts)
	return counts
}
