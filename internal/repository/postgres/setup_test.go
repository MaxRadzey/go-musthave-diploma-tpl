package postgres_test

import (
	"context"
	"testing"

	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/repository/postgres"
)

// testRepos — все репозитории для тестов. Один setup, один TRUNCATE.
type testRepos struct {
	User  *postgres.UserRepository
	Order *postgres.OrderRepository
}

// setupDB очищает таблицы orders и users и возвращает все репозитории.
func setupDB(t *testing.T) *testRepos {
	t.Helper()
	_, err := testDB.ExecContext(context.Background(), "TRUNCATE orders, users RESTART IDENTITY CASCADE")
	if err != nil {
		t.Fatalf("truncate: %v", err)
	}
	return &testRepos{
		User:  postgres.NewUserRepository(testDB),
		Order: postgres.NewOrderRepository(testDB),
	}
}
