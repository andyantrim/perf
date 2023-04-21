package handler_test

import (
	"testing"

	"github.com/andyantrim/perf"
	"github.com/andyantrim/perf/handler"
)

func BenchmarkProcessMessage(b *testing.B) {
	fn := func(msg perf.Output) {}
	processor := handler.NewProcessor(fn)
	// Start the handler
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
		processor.Process(message)
	}
}
