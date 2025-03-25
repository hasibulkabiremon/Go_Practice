package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	// ?
	a := []float64 {}
	for _,x := range costs{
		if x.day == day{
			a = append(a, x.value)
		}
	}
	return a
}
