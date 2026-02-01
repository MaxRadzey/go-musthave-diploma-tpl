package repository

import "fmt"

// ErrDuplicateLogin — логин уже занят.
type ErrDuplicateLogin struct {
	Login string
}

func (e *ErrDuplicateLogin) Error() string {
	return fmt.Sprintf("login %q already exists", e.Login)
}

// ErrUserNotFound — пользователь не найден.
type ErrUserNotFound struct {
	Login string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("user with login %q not found", e.Login)
}

// ErrOrderNotFound — заказ с таким номером не найден.
type ErrOrderNotFound struct {
	Number string
}

func (e *ErrOrderNotFound) Error() string {
	return fmt.Sprintf("order with number %q not found", e.Number)
}
