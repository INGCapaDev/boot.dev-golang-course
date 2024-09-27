package main

import "time"

func processMessages(messages []string) []string {
	processedMsgs := make([]string, len(messages))
	ch := make(chan string, len(messages))

	for _, msg := range messages {
		go func(m string) {
			ch <- process(m)
		}(msg)
	}

	for i := 0; i < len(messages); i++ {
		processedMsgs[i] = <-ch
	}

	return processedMsgs
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
