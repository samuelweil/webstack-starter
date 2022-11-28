package task

type TaskService interface {
	AllTasks() []Task
}

type inMemoryRepo struct {
	tasks []Task
}

func (r *inMemoryRepo) AllTasks() []Task {
	return r.tasks
}

func NewService() TaskService {
	return &inMemoryRepo{
		tasks: make([]Task, 0),
	}
}
