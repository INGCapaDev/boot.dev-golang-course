package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analyticsPtr *Analytics, msg Message) {
	analyticsPtr.MessagesTotal++
	if msg.Success {
		analyticsPtr.MessagesSucceeded++
		return
	}
	analyticsPtr.MessagesFailed++
}
