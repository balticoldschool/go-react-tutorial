package tasks

import "github.com/balticoldschool/go-react-tutorial/src/utils/errors"

type Service interface {
	AddTask(task *Task) *Task
	GetAll() *[]*Task
	GetById(id uint) (*Task, *errors.RestErr)
	UpdateTaskCompletion(id uint) (*Task, *errors.RestErr)
	RemoveTask(id uint) (*Task, *errors.RestErr)
}

type service struct{}

func NewTaskService() Service {
	return &service{}
}

func (s service) AddTask(task *Task) *Task {
	return addToTasks(task)
}

func (s service) GetAll() *[]*Task {
	return &TaskStorage
}

func (s service) GetById(id uint) (*Task, *errors.RestErr) {
	task := findById(id)

	if len(TaskStorage) == 0 || task == nil {
		return task, errors.NewNotFoundErr("no task found")
	}

	return task, nil
}

func (s service) UpdateTaskCompletion(id uint) (*Task, *errors.RestErr) {
	return updateCompletion(id)
}

func (s service) RemoveTask(id uint) (*Task, *errors.RestErr) {
	return removeById(id)
}
