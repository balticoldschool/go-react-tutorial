package tasks

type Service interface {
	AddTask(task *Task) *Task
	GetAll() *[]Task
}

type service struct{}

func NewTaskService() Service {
	return &service{}
}

func (s service) AddTask(task *Task) *Task {
	newTask := *task
	newTask.ID = len(TaskStorage) + 1

	TaskStorage = append(TaskStorage, newTask)

	return &newTask
}

func (s service) GetAll() *[]Task {
	tasks := TaskStorage
	return &tasks
}
