package load_test

import (
	"testing"
	"github.com/andyantrim/perf/load"
	"github.com/spf13/viper"
)

func BenchmarkV1(b *testing.B) {
	tests := []struct {
		name string
		toUsers []string
		fromUser int
		content string
	}{
		{
			name: "TestProcessMessage",
			toUsers: []string{"1", "2", "3", "4", "5"},
			fromUser: 1,
			content: "Hello World!",
		},
	}
	viper.Set("debug", false)
	for _, tt := range tests {
		msg := load.Message{
			UUID: "1234567890",
			SenderID: tt.fromUser,
			Recievers: tt.toUsers,
			Content: tt.content,
		}
		for i := 0; i < b.N; i++ {
	 		load.ProcessMessage(msg)
		}
	}
}


func BenchmarkV2(b *testing.B) {
	tests := []struct {
		name string
		toUsers []string
		fromUser int
		content string
	}{
		{
			name: "TestProcessMessage",
			toUsers: []string{"1", "2", "3", "4", "5"},
			fromUser: 1,
			content: "Hello World!",
		},
	}
	viper.Set("debug", false)
	for _, tt := range tests {
		msg := load.Message{
			UUID: "1234567890",
			SenderID: tt.fromUser,
			Recievers: tt.toUsers,
			Content: tt.content,
		}
		for i := 0; i < b.N; i++ {
	 		load.ProcessMessageV2(msg, viper.GetBool("debug"))
		}
	}
}
