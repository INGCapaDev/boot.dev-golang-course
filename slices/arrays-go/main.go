package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	pCost := len(primary)
	sCost := pCost + len(secondary)
	tCost := sCost + len(tertiary)
	return messages, [3]int{pCost, sCost, tCost}
}
