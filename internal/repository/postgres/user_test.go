package postgres_test

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/repository"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/repository/postgres"
	"github.com/MaxRadzey/go-musthave-diploma-tpl/internal/storage"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	testDSN string
	testDB  *sql.DB
)

// defaultTestDSN — тот же DSN, что и у контейнера по умолчанию (config.DatabaseDSN).
const defaultTestDSN = "postgres://shortener:shortener@localhost:5432/shortener?sslmode=disable"

func TestMain(m *testing.M) {
	testDSN = os.Getenv("TEST_DATABASE_DSN")
	if testDSN == "" {
		testDSN = defaultTestDSN
	}

	if err := storage.RunMigrations(testDSN, "../../../migrations"); err != nil {
		os.Stderr.WriteString("migrations: " + err.Error() + "\n")
		os.Exit(1)
	}

	var err error
	testDB, err = sql.Open("pgx", testDSN)
	if err != nil {
		os.Stderr.WriteString("open db: " + err.Error() + "\n")
		os.Exit(1)
	}
	defer testDB.Close()

	os.Exit(m.Run())
}

func setupDB(t *testing.T) (*sql.DB, *postgres.UserRepository) {
	_, err := testDB.ExecContext(context.Background(), "TRUNCATE users RESTART IDENTITY CASCADE")
	if err != nil {
		t.Fatalf("truncate: %v", err)
	}
	repo := postgres.NewUserRepository(testDB)
	return testDB, repo
}

func TestUserRepository_Create(t *testing.T) {
	_, repo := setupDB(t)
	ctx := context.Background()

	user, err := repo.Create(ctx, "alice", "hash123")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	if user.ID == 0 {
		t.Error("expected non-zero ID")
	}
	if user.Login != "alice" {
		t.Errorf("Login: got %q", user.Login)
	}
	if user.PasswordHash != "hash123" {
		t.Errorf("PasswordHash: got %q", user.PasswordHash)
	}
	if !user.Active {
		t.Error("expected Active true")
	}
	if user.CreatedAt.IsZero() {
		t.Error("expected non-zero CreatedAt")
	}
}

func TestUserRepository_Create_DuplicateLogin(t *testing.T) {
	_, repo := setupDB(t)
	ctx := context.Background()

	_, err := repo.Create(ctx, "Max", "hash1")
	if err != nil {
		t.Fatalf("first Create: %v", err)
	}

	_, err = repo.Create(ctx, "Max", "hash2")
	if err == nil {
		t.Fatal("expected ErrDuplicateLogin")
	}
	var dup *repository.ErrDuplicateLogin
	if !errors.As(err, &dup) {
		t.Fatalf("expected *ErrDuplicateLogin, got %T: %v", err, err)
	}
	if dup.Login != "Max" {
		t.Errorf("ErrDuplicateLogin.Login: got %q", dup.Login)
	}
}

func TestUserRepository_GetByLogin_Found(t *testing.T) {
	_, repo := setupDB(t)
	ctx := context.Background()

	created, err := repo.Create(ctx, "Max", "secret")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}

	user, err := repo.GetByLogin(ctx, "Max")
	if err != nil {
		t.Fatalf("GetByLogin: %v", err)
	}
	if user.ID != created.ID {
		t.Errorf("ID: got %d", user.ID)
	}
	if user.Login != "Max" {
		t.Errorf("Login: got %q", user.Login)
	}
	if user.PasswordHash != "secret" {
		t.Errorf("PasswordHash: got %q", user.PasswordHash)
	}
}

func TestUserRepository_GetByLogin_NotFound(t *testing.T) {
	_, repo := setupDB(t)
	ctx := context.Background()

	_, err := repo.GetByLogin(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected ErrUserNotFound")
	}
	var notFound *repository.ErrUserNotFound
	if !errors.As(err, &notFound) {
		t.Fatalf("expected *ErrUserNotFound, got %T: %v", err, err)
	}
}
