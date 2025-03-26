package tasks

var (
	TaskStorage = make([]Task, 0)
)

type Task struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}
