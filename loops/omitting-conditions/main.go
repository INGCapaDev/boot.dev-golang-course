package main

func maxMessages(thresh int) int {
	baseMsgCost := 100
	totalMessagesCost := 0
	for i := 0; ; i++ {
		totalMessagesCost += baseMsgCost + i
		if thresh < totalMessagesCost {
			return i
		}
	}
}
