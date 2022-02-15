package main

import "fmt"

type response struct {
	message string
}

func newResponse(message string) *response {
	response := response{message: message}
	return &response
}

func messageSeq() func() *response {
	i := 0
	return func() *response {
		i++
		return newResponse(fmt.Sprintf("Ola mundo %d", i))
	}
}

func main() {
	messageNext := messageSeq()

	messages := make(chan response)

	go func() { messages <- *messageNext() }()

	msg := <-messages
	fmt.Println(msg.message)
}
