package main

import (
	"time"
)

func processMessages(messages []string) []string {
	processedMessages := make([]string, 0)
	ch := make(chan string, len(messages))
	go func() {
		for _, message := range messages {
			ch <- process(message)
		}
		close(ch)
	}()
	for processedMessage := range ch {
		processedMessages = append(processedMessages, processedMessage)
	}
	return processedMessages
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
