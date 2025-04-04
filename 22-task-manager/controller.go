package main

import (
	"errors"

	"example.com/task-manager/models"
	"github.com/google/uuid"
)

func GetTasks(tasksList []models.Task) any {
	if len(tasksList) > 0 {
		return tasksList
	}

	return EmptyTasksMessage
}

func AddTask(tasksList []models.Task, title string, description string, status string) ([]models.Task, string) {
	var newTask models.Task = models.NewTask(title, description, status)
	tasksList = append(tasksList, newTask)
	return tasksList, TaskCreatedMessage
}

func UpdateTask(tasksList []models.Task, taskId uuid.UUID, param string, value string) ([]models.Task, error) {
	for i := range tasksList {
		if tasksList[i].Id == taskId {
			switch param {
			case "Title":
				tasksList[i].Title = value
			case "Description":
				tasksList[i].Description = value
			case "Status":
				tasksList[i].Status = value
			default:
				return nil, errors.New(UnsupportedParamMesssage)
			}
			break
		}
	}

	return tasksList, nil
}

func FindTaskFromTasklist(taskList []models.Task, taskId uuid.UUID) ([]models.Task, error) {
	for _, task := range taskList {
		if task.Id == taskId {
			return []models.Task{task}, nil
		}
	}

	return nil, errors.New("task was not found")
}

func DeleteTaskFromTaskList(tasksList []models.Task, taskId uuid.UUID) ([]models.Task, error) {
	for index := range tasksList {
		if tasksList[index].Id == taskId {
			tasksList = append(tasksList[:index], tasksList[index+1:]...)
			return tasksList, nil
		}
	}

	return nil, errors.New(InvalidTaskId)
}
