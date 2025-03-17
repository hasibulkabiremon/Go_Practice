package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?
	for i, a := range msg {
		for _, b := range badWords {
			if a == b {
				return i
			}
		}
	}
	return -1
} 