package models

import "github.com/google/uuid"

type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	Status      string
}

func NewTask(title, description, status string) Task {
	return Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      status,
	}
}
