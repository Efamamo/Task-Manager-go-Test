package usecases

import (
	"errors"
	"strings"

	domain "github.com/Task-Management-go/Domain"
	repository "github.com/Task-Management-go/Repository"
)



var taskRepository TaskInterface = &repository.TaskRepository{}

func GetTasks() (*[]domain.Task, error) {

	tasks, err := taskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil

}

func GetTaskByID(id string) (*domain.Task, error) {
	task, err := taskRepository.FindOne(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func UpdateItem(ID string, updatedTask domain.Task) (*domain.Task, error) {

	utask, err := taskRepository.UpdateOne(ID, updatedTask)

	if err != nil {
		return nil, err
	}

	return utask, nil

}

func DeleteTask(ID string) (*domain.Task, error) {

	task, err := taskRepository.DeleteOne(ID)

	if err != nil {
		return nil, err
	}
	return task, nil

}

func AddTask(task domain.Task) (*domain.Task, error) {
	if strings.ToLower(task.Status) != "in progress" && strings.ToLower(task.Status) != "completed" && strings.ToLower(task.Status) != "pending" {
		return nil, errors.New("status error")
	}

	t, err := taskRepository.Save(task)

	if err != nil {
		return nil, err
	}

	return t, nil
}
