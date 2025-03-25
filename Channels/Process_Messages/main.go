package main

import "time"

func processMessages(messages []string) []string {
	ch := make(chan string)
	result := []string{}
	for _, message := range messages {
		go func(msg string) {
			ch <- process(msg)
		}(message)
	}

	for range messages {
		result = append(result, <-ch)
	}
	
	return result
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}