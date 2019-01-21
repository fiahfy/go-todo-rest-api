package usecase

import (
	"github.com/fiahfy/go-todo-rest-api/domain/model"
	"github.com/fiahfy/go-todo-rest-api/domain/repository"
)

type TodoUseCase interface {
	Find(int) (*model.Todo, error)
	FindAll() ([]*model.Todo, error)
}

type todoUseCase struct {
	r repository.TodoRepository
}

func NewTodoUseCase(r repository.TodoRepository) TodoUseCase {
	return &todoUseCase{r}
}

func (u *todoUseCase) Find(id int) (*model.Todo, error) {
	return u.r.Find(id)
}

func (u *todoUseCase) FindAll() ([]*model.Todo, error) {
	return u.r.FindAll()
}