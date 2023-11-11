package main

import (
	"fmt"
	"fr/config"
	"fr/internal/app"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/mattes/migrate/source/file"
	"log"
)

func migrateUp(cfg *config.Config) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Username, cfg.PG.Password, cfg.PG.Host, cfg.PG.Port, cfg.PG.Name)
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatal(err)
	}
}
func init() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	migrateUp(cfg)

	fmt.Println("migrated")

}
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
