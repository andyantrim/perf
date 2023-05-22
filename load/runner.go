package load

import (
	"fmt"
	"log"
	"strconv"
	"github.com/andyantrim/perf/fib"
	"github.com/spf13/viper"
)

var Debug bool = true

type Message struct {
	UUID string
	SenderID int
	Recievers []string
	Content string
}

func Do(m Message, to int, _ string) {
	var val int
	for i := range m.Recievers {
		n, err := strconv.Atoi(m.Recievers[i])
		if err != nil {
			continue
		}
		val = fib.FibRecursive(n^2)
		val = ((n^2)*(val^n))^val
	}

	for val > 0 {
		val--
	}

	return
}

func ProcessMessage(message Message) {
	// Do something
	for _, reciever := range message.Recievers {
		// Send message to reciever
		userID, err := strconv.Atoi(reciever)
		if err != nil {
			log.Printf("Failed to convert reciever %s to int\n", reciever)
			continue
		}
		// If debug mode is on, print the message info
		entry := fmt.Sprintf("Sending message %s to user %d\n", message.UUID, userID)
		if viper.GetBool("debug") {
			log.Println(entry)
		}
		Do(message, userID, entry)
	}
}

func ProcessMessageV2(message Message, debug bool) {
	// Do something
	for _, reciever := range message.Recievers {
		// Send message to reciever
		userID, err := strconv.Atoi(reciever)
		if err != nil {
			log.Println("Failed to convert reciever" + reciever + " to int")
			continue
		}
		// If debug mode is on, print the message info
		entry := "Sending message" + message.UUID + " to user " + reciever
		if debug {
			log.Println(entry)
		}
		Do(message, userID, entry)
	}
}
