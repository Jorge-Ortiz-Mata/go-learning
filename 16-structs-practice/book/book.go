package book

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const bookFile = "book.json"
const bookContentMinimumLenth int = 20

type Book struct {
	Title     string    `json:"title"`      // tags to format the JSON response
	Content   string    `json:"content"`    // tags to format the JSON response
	CreatedAt time.Time `json:"created_at"` // tags to format the JSON response
}

func New(title, content string) *Book {
	return &Book{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
}

func (b Book) titleValidations() (bool, error) {
	if b.Title == "" {
		return false, errors.New("the title must not be blank")
	}

	return true, nil
}

func (b Book) contentValidations() (bool, error) {
	if b.Content == "" {
		return false, errors.New("the content must not be blank")
	}

	if len(b.Content) < bookContentMinimumLenth {
		return false, errors.New("the content must be at least 20 characters")
	}

	return true, nil
}

func (b Book) Valid() (bool, error) {
	titleIsValid, title_err := b.titleValidations()
	contentIsValid, content_err := b.contentValidations()

	if !titleIsValid {
		return false, title_err
	} else if !contentIsValid {
		return false, content_err
	} else {
		return true, nil
	}
}

func (b Book) Export() {
	// // ===== One way of doing it =====
	// contentToExport := fmt.Sprintf(`{"title": "%v", "content": "%v", "created_at": "%v"}`, b.Title, b.Content, b.CreatedAt)
	// os.WriteFile(bookFile, []byte(contentToExport), 0644)
	// fmt.Println("The book has been exported!")
	// ===================================

	// ===== Second way of doing it. =====
	json, err := json.Marshal(b) // Converts a struct into a JSON object
	if err != nil {
		fmt.Println(err)
	}
	os.WriteFile(bookFile, json, 0644)
	// ===================================
}

func (b Book) Call() {
	response, err := b.Valid()

	if response {
		b.Export()
	} else {
		fmt.Println(err)
	}
}
