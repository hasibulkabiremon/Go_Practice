package main

func maxMessages(thresh int) int {
	var msg int
	for i := 100; i <= thresh; i += msg + 100 {
		msg++
		// println(msg, ":", i)
	}
	return msg
}
