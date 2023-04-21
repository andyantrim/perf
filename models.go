package perf

type Message struct {
	Description string
	FromUser    User
	ToList      []User
	EntityType  string
	ActionType  string
	EntityID    int
	EntityName  string
}

type User struct {
	ID   int
	Name string
}

type Output struct {
	Title       string
	Description string
	UserID      int
}

type Processor interface {
	Process(msg Message) error
}
