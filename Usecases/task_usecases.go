package usecases

import (
	"errors"
	"strings"

	domain "github.com/Task-Management-go/Domain"
)

type TaskService struct {
	TaskRepo TaskInterface
}

func (ts *TaskService) GetTasks() (*[]domain.Task, error) {

	tasks, err := ts.TaskRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil

}

func (ts *TaskService) GetTaskByID(id string) (*domain.Task, error) {
	task, err := ts.TaskRepo.FindOne(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *TaskService) UpdateItem(ID string, updatedTask domain.Task) error {

	if strings.ToLower(updatedTask.Status) != "in progress" && strings.ToLower(updatedTask.Status) != "completed" && strings.ToLower(updatedTask.Status) != "pending" {
		return errors.New("status error")
	}

	err := ts.TaskRepo.UpdateOne(ID, updatedTask)

	if err != nil {
		return err
	}

	return nil

}

func (ts *TaskService) DeleteTask(ID string) (*domain.Task, error) {

	task, err := ts.TaskRepo.DeleteOne(ID)

	if err != nil {
		return nil, err
	}
	return task, nil

}

func (ts *TaskService) AddTask(task domain.Task) (*domain.Task, error) {
	if strings.ToLower(task.Status) != "in progress" && strings.ToLower(task.Status) != "completed" && strings.ToLower(task.Status) != "pending" {
		return nil, errors.New("status error")
	}

	t, err := ts.TaskRepo.Save(task)

	if err != nil {
		return nil, err
	}

	return t, nil
}
