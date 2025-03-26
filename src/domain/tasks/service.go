package tasks

import "github.com/balticoldschool/go-react-tutorial/src/utils/errors"

type Service interface {
	AddTask(task *Task) *Task
	GetAll() *[]Task
	GetById(id uint) (*Task, *errors.RestErr)
}

type service struct{}

func NewTaskService() Service {
	return &service{}
}

func (s service) AddTask(task *Task) *Task {
	newTask := *task
	newTask.ID = uint(len(TaskStorage) + 1)

	TaskStorage = append(TaskStorage, newTask)

	return &newTask
}

func (s service) GetAll() *[]Task {
	tasks := TaskStorage
	return &tasks
}

func (s service) GetById(id uint) (*Task, *errors.RestErr) {
	task := findById(id)

	if len(TaskStorage) == 0 || task == nil {
		return task, errors.NewNotFoundErr("no task found")
	}

	return task, nil
}
