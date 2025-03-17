package main

func sum(nums ...int) int {
	// ?
	var s int
	for x:=range nums {
		s+=nums[x]
	}
	return s
}
