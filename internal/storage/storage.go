package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/repository"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/repository/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Storage — подключение к БД и репозитории.
type Storage struct {
	db                  *sql.DB
	UserRepository      repository.UserRepository
	OrderRepository     repository.OrderRepository
	WithdrawalRepository repository.WithdrawalRepository
}

// InitializeStorage запускает миграции, подключается к БД и создаёт репозитории.
// При пустом dsn возвращает ошибку.
func InitializeStorage(dsn string) (*Storage, error) {
	if dsn == "" {
		return nil, errors.New("database DSN required")
	}
	if err := RunMigrations(dsn, ""); err != nil {
		return nil, fmt.Errorf("migrations: %w", err)
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return &Storage{
		db:                   db,
		UserRepository:       postgres.NewUserRepository(db),
		OrderRepository:      postgres.NewOrderRepository(db),
		WithdrawalRepository: postgres.NewWithdrawalRepository(db),
	}, nil
}

// Close закрывает подключение к БД.
func (s *Storage) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
