package main

import (
	"testing"

	"example.com/task-manager/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGetTasksMessageIfTheGetTasksListIsEmpty(t *testing.T) {
	var tasksList []models.Task
	result := GetTasks(tasksList)

	assert.Equal(t, result, "There is no tasks added", "The message must be equal")
}

func TestGetTasksAfterAddingOne(t *testing.T) {
	var tasksList []models.Task
	tasksListUpdated, _ := AddTask(tasksList, "Send email", "Open gmail to send party invitations", "completed")
	result := GetTasks(tasksListUpdated)

	assert.Equal(t, len(tasksListUpdated), 1, "The task has been added to the list")
	assert.NotEqual(t, result, "There is no tasks added", "The message must be equal")
}

func TestAddTaskToTasksList(t *testing.T) {
	var tasksList []models.Task
	tasksList, result := AddTask(tasksList, "Buy eggs", "Go to the store and purchase eggs", "pending")

	assert.Equal(t, tasksList[0].Title, "Buy eggs", "The title must match")
	assert.Equal(t, result, "The task has been added to the list", "The message must be equal")
	assert.Equal(t, len(tasksList), 1, "The task has been added to the list")
}

func TestUpdateTaskWithoutValidParam(t *testing.T) {
	var tasksList []models.Task
	tasksList, _ = AddTask(tasksList, "Send email", "Open gmail to send party invitations", "pending")
	tasksList, _ = AddTask(tasksList, "Buy Eggs", "Go to the supermarket to buy eggs", "pending")

	assert.Equal(t, tasksList[0].Status, "pending", "The status must be set")
	assert.Equal(t, len(tasksList), 2, "The task has been added to the list")

	_, err := UpdateTask(tasksList, tasksList[1].Id, "status", "completed")

	assert.Equal(t, err.Error(), UnsupportedParamMesssage, "The param to update is invalid")
}

func TestUpdateTaskWithValidParam(t *testing.T) {
	var tasksList []models.Task
	tasksList, _ = AddTask(tasksList, "Send email", "Open gmail to send party invitations", "pending")
	tasksList, _ = AddTask(tasksList, "Buy Eggs", "Go to the supermarket to buy eggs", "pending")

	assert.Equal(t, tasksList[0].Status, "pending", "The status must be set")
	assert.Equal(t, len(tasksList), 2, "The task has been added to the list")

	result, _ := UpdateTask(tasksList, tasksList[0].Id, "Status", "completed")

	assert.Equal(t, result[0].Status, "completed", "The status has been updated")
	assert.Equal(t, len(result), 2, "The task has been added to the list")
}

func TestDeleteTaskWithInvalidId(t *testing.T) {
	var tasksList []models.Task
	tasksList, _ = AddTask(tasksList, "Send email", "Open gmail to send party invitations", "pending")
	tasksList, _ = AddTask(tasksList, "Buy Eggs", "Go to the supermarket to buy eggs", "pending")
	tasksList, _ = AddTask(tasksList, "Purchase a TV", "Purchase a new TV for the living room", "completed")

	var taskId uuid.UUID = uuid.New()

	_, err := DeleteTaskFromTaskList(tasksList, taskId)

	assert.Equal(t, err.Error(), InvalidTaskId, "Invalid task id")
}

func TestDeleteTaskWithValidId(t *testing.T) {
	var tasksList []models.Task
	tasksList, _ = AddTask(tasksList, "Send email", "Open gmail to send party invitations", "pending")
	tasksList, _ = AddTask(tasksList, "Buy Eggs", "Go to the supermarket to buy eggs", "pending")
	tasksList, _ = AddTask(tasksList, "Purchase a TV", "Purchase a new TV for the living room", "completed")

	var taskId uuid.UUID = tasksList[1].Id

	result, err := FindTaskFromTasklist(tasksList, taskId)

	assert.Nil(t, err, "The error must be nil")
	assert.Equal(t, result[0].Title, "Buy Eggs", "The title must be present")

	tasksList, _ = DeleteTaskFromTaskList(tasksList, taskId)
	result, err = FindTaskFromTasklist(tasksList, taskId)

	assert.Equal(t, err.Error(), "task was not found")
	assert.Nil(t, result, nil, "The result must be empty")
}
