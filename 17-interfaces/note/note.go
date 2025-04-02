package note

import (
	"fmt"
	"time"
)

type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

func New(title, content string) Note {
	return Note{
		Title:   title,
		Content: content,
	}
}

func (n Note) Display() {
	fmt.Printf("This is the title: %v; and this is the content: %v\n", n.Title, n.Content)
}

func (n Note) Save() string {
	return fmt.Sprintf("Content to save - Title: %v; Content: %v", n.Title, n.Content)
}
