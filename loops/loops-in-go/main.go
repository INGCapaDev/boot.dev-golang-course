package main

func bulkSend(numMessages int) float64 {
	baseCost := 1.0
	baseAdditionalFee := 0.01
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {

		totalCost += baseCost + (float64(i) * baseAdditionalFee)
	}
	return totalCost
}
