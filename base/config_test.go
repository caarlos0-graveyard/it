package base

import (
	"os"
	"testing"

	"github.com/caarlos0/env"
	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	createTestDB := os.Getenv("DROP_TEST_DB")
	os.Unsetenv("DROP_TEST_DB")
	defer os.Setenv("DROP_TEST_DB", createTestDB)

	postgresURL := os.Getenv("POSTGRES_URL")
	os.Unsetenv("POSTGRES_URL")
	defer os.Setenv("POSTGRES_URL", postgresURL)

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	os.Unsetenv("MIGRATIONS_PATH")
	defer os.Setenv("MIGRATIONS_PATH", migrationsPath)

	var cfg Config
	env.Parse(&cfg)

	assert.Equal(t, true, cfg.DropDB)
	assert.Equal(t, "postgres://localhost:5432", cfg.PostgresURL)
	assert.Equal(t, "./migrations", cfg.MigrationsFolder)
}
