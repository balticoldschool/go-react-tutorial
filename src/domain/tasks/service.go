package tasks

type Service interface {
	AddTask(task *Task) *Task
}

type service struct{}

func NewTaskService() Service {
	return &service{}
}

func (s service) AddTask(task *Task) *Task {
	newTask := *task
	newTask.ID = len(TaskStorage)

	TaskStorage = append(TaskStorage, newTask)

	return &newTask
}
