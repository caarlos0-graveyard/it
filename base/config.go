package base

// Config for ITs
type Config struct {
	CreateDB         bool   `env:"CREATE_TEST_DB" envDefault:"true"`
	MigrateDB        bool   `env:"MIGRATE_DB" envDefault:"true"`
	PostgresURL      string `env:"POSTGRES_URL" envDefault:"postgres://localhost:5432"`
	DatabaseName     string
	MigrationsFolder string `env:"MIGRATIONS_PATH" envDefailt:"./migrations"`
	SSLEnabled       bool
}
