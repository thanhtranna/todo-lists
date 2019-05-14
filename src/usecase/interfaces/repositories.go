package interfaces

import "todo-lists/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindByName(string) ([]domain.User, error)
	FindAll() ([]domain.User, error)
}