package service

import "fmt"

// ErrInvalidCredentials — неверная пара логин/пароль.
type ErrInvalidCredentials struct {
	Login string
}

func (e *ErrInvalidCredentials) Error() string {
	return fmt.Sprintf("invalid credentials for login %q", e.Login)
}

// ErrValidation — ошибка валидации (пустой логин/пароль и т.д.).
type ErrValidation struct {
	Msg string
}

func (e *ErrValidation) Error() string {
	return e.Msg
}
