package handler

import (
	"fmt"

	"github.com/andyantrim/perf"
)

type Processor struct {
	Input  chan perf.Message
	Output chan perf.Output
}

// Creates a new Processor with the channels provided
func NewProcessor(i chan perf.Message, o chan perf.Output) *Processor {
	return &Processor{i, o}
}

// Starts the Processor
func (p *Processor) Listen() {
	for {
		select {
		case msg := <-p.Input:
			p.ProcessMessage(msg)
		}
	}
}

func (p *Processor) ProcessMessage(msg perf.Message) {
	// Make a title from the message
	title := fmt.Sprintf("%s %sed a %s", msg.FromUser.Name, msg.ActionType, msg.EntityName)

	// Clean the userIDlist
	userIDList := make([]int, 0)
	for _, user := range msg.ToList {
		if user.ID != msg.FromUser.ID && user.ID != 0 {
			userIDList = append(userIDList, user.ID)
		}
	}

	// Create an output for each user
	for _, userID := range userIDList {
		output := perf.Output{
			Title:       title,
			Description: msg.Description,
			UserID:      userID,
		}
		p.Output <- output
	}

}
