package it

import (
	"bufio"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
)

// DBPoolFn blah
type DBPoolFn func(url string) *sqlx.DB

// DB blah
type DB struct {
	connect DBPoolFn
	cfg     Config
	con     *sqlx.DB
}

// Init blah
func (db *DB) Init() *sqlx.DB {
	var dbURL string
	if db.cfg.CreateDB {
		dbURL = db.cfg.PostgresURL + createTestDatabase(db.cfg)
	}
	db.con = prepareTestDB(db.connect(dbURL), db.cfg)
	return db.con
}

// Shutdown blah
func (db *DB) Shutdown() {
	db.con.Close()
	if db.cfg.CreateDB {
		pgExec("DROP DATABASE "+db.cfg.DatabaseName, db.cfg)
	}
}

func createTestDatabase(cfg Config) string {
	name := randomStr()
	log.Println("Create-ing test database " + name)
	pgExec("CREATE DATABASE "+name, cfg)
	return name
}

func prepareTestDB(db *sqlx.DB, cfg Config) *sqlx.DB {
	if cfg.MigrateDB {
		migrate(db, cfg)
	}
	return db
}

func migrate(db *sqlx.DB, cfg Config) {
	log.Println("Migrate-ing database...")
	files, _ := filepath.Glob(filepath.Join(cfg.MigrationsFolder, "*.sql"))
	for _, file := range files {
		file, _ := os.Open(file)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var sql string
		for scanner.Scan() {
			sql += scanner.Text()
		}
		if _, err := db.Exec(sql); err != nil {
			panic(err)
		}
	}
}

func pgExec(stm string, cfg Config) {
	db, err := sqlx.Connect("postgres", cfg.PostgresURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if _, err = db.Exec(stm); err != nil {
		panic(err)
	}
}
