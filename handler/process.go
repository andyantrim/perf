package handler

import (
	"fmt"

	"github.com/andyantrim/perf"
)

type Processor struct {
	SendFn func(perf.Output)
}

// Creates a new Processor with the channels provided
func NewProcessor(sendFn func(perf.Output)) *Processor {
	return &Processor{
		SendFn: sendFn,
	}
}

func (p *Processor) Process(msg perf.Message) error {
	// Make a title from the message
	title := fmt.Sprintf("%s %sed a %s", msg.FromUser.Name, msg.ActionType, msg.EntityName)
	// Clean the userIDlist
	for _, user := range msg.ToList {
		if user.ID != msg.FromUser.ID && user.ID != 0 {
			output := perf.Output{
				Title:       title,
				Description: msg.Description,
				UserID:      user.ID,
			}
			p.SendFn(output)
		}
	}

	return nil
}
