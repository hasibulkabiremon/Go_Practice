package main

func bulkSend(numMessages int) float64 {

	var total int
	for i := 0; i < numMessages ; i++{
		total += 100 + i 
	}
	return float64(total)/100
}