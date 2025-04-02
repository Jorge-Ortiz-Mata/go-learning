package todo

import (
	"fmt"
)

type Todo struct {
	Text string
}

func New(text string) Todo {
	return Todo{
		Text: text,
	}
}

func (t Todo) Display() {
	fmt.Printf("This is the TODO content: %v\n", t.Text)
}

func (t Todo) Save() string {
	return fmt.Sprintf("Content to save - Text: %v\n", t.Text)
}
