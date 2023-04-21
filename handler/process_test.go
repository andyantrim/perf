package handler_test

import (
	"testing"

	"github.com/andyantrim/perf"
	"github.com/andyantrim/perf/handler"
)

func BenchmarkProcessMessage(b *testing.B) {
	// Create the channels and handler
	input := make(chan perf.Message)
	output := make(chan perf.Output)
	processor := handler.NewProcessor(input, output)
	outputMessage := perf.Output{}
	// Start the handler
	go processor.Listen()
	go func() {
		for {
			select {
			case msg := <-output:
				outputMessage = msg
				_ = outputMessage
			}
		}
	}()
	message := perf.Message{
		Description: "This is a test message",
		FromUser: perf.User{
			ID:   1,
			Name: "Andy",
		},
		ToList: []perf.User{
			{
				ID:   2,
				Name: "Bob",
			},
			{
				ID:   3,
				Name: "Charlie",
			},
			{
				ID:   4,
				Name: "Diane",
			},
		},
		EntityType: "post",
		ActionType: "like",
		EntityID:   1,
		EntityName: "post",
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		// Send the message to the handler
		input <- message
	}
}
