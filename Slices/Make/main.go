package main

func getMessageCosts(messages []string) []float64{
    // var cost []float64
    println("----------")
    cost := make([]float64,len(messages))
    println(cap(cost))
    for i, x := range messages{
        cost[i] = float64(len(x))* 0.01
        println(cap(cost))
    }
    println(cap(cost))
    return cost
}