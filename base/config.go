package base

// Config for ITs
type Config struct {
	DropDB           bool   `env:"DROP_TEST_DB" envDefault:"true"`
	PostgresURL      string `env:"POSTGRES_URL" envDefault:"postgres://localhost:5432?sslmode=disable"`
	DatabaseName     string
	MigrationsFolder string `env:"MIGRATIONS_PATH" envDefault:"./migrations"`
}
