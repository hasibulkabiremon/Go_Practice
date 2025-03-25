package main

func countReports(numSentCh chan int) int {
	totalReports := 0
	for {
		numReports, ok := <-numSentCh
		if !ok {
			break
		}
		totalReports += numReports
	}
	return totalReports
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}