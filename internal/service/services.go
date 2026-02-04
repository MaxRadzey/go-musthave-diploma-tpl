package service

import "github.com/MaxRadzey/go-musthave-diploma-tpl/internal/storage"

// Services — контейнер всех сервисов приложения.
type Services struct {
	User    *UserService
	Order   *OrderService
	Balance *BalanceService
}

// NewServices создаёт все сервисы из Storage (все репозитории берутся из него).
func NewServices(st *storage.Storage) *Services {
	return &Services{
		User:    NewUserService(st.UserRepository),
		Order:   NewOrderService(st.OrderRepository),
		Balance: NewBalanceService(st.OrderRepository, st.WithdrawalRepository),
	}
}
