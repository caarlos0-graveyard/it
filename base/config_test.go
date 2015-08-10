package base

import (
	"os"
	"testing"

	"github.com/caarlos0/env"
	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	createTestDB := os.Getenv("CREATE_TEST_DB")
	os.Unsetenv("CREATE_TEST_DB")
	defer os.Setenv("CREATE_TEST_DB", createTestDB)
	migrateDB := os.Getenv("MIGRATE_DB")
	os.Unsetenv("MIGRATE_DB")
	defer os.Setenv("MIGRATE_DB", migrateDB)
	postgresURL := os.Getenv("POSTGRES_URL")
	os.Unsetenv("POSTGRES_URL")
	defer os.Setenv("POSTGRES_URL", postgresURL)
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	os.Unsetenv("MIGRATIONS_PATH")
	defer os.Setenv("MIGRATIONS_PATH", migrationsPath)

	var cfg Config
	env.Parse(&cfg)

	assert.Equal(t, true, cfg.CreateDB)
	assert.Equal(t, true, cfg.MigrateDB)
	assert.Equal(t, "postgres://localhost:5432", cfg.PostgresURL)
	assert.Equal(t, "./migrations", cfg.MigrationsFolder)
}
