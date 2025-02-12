package main

const (
	costPerCharacter = 0.01
)

func getMessageCosts(messages []string) []float64 {
	messageCosts := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
		messageCosts[i] = float64(len(messages[i])) * costPerCharacter
	}
	return messageCosts
}
