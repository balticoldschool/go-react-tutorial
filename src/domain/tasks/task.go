package tasks

import (
	"github.com/balticoldschool/go-react-tutorial/src/utils/errors"
)

var (
	TaskStorage = make([]*Task, 0)
)

type Task struct {
	ID        uint   `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func addToTasks(task *Task) *Task {
	task.ID = uint(len(TaskStorage) + 1)
	TaskStorage = append(TaskStorage, task)

	return task
}

func findById(id uint) *Task {
	for i := range TaskStorage {
		if TaskStorage[i].ID == id {
			return TaskStorage[i]
		}
	}

	return nil
}

func updateCompletion(id uint) (*Task, *errors.RestErr) {
	task := findById(id)
	if task == nil {
		return nil, errors.NewNotFoundErr("no task found")
	}

	task.Completed = !task.Completed

	return task, nil
}
