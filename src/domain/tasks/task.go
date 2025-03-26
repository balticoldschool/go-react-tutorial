package tasks

var (
	TaskStorage = make([]Task, 0)
)

type Task struct {
	ID        uint   `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func findById(id uint) *Task {
	for _, task := range TaskStorage {
		if task.ID == id {
			return &task
		}
	}
	return nil
}
