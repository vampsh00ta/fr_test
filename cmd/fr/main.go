package main

import (
	"fr/config"
	_ "fr/docs"
	"fr/internal/app"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

//	func migrateUp(cfg *config.Config) {
//		url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
//			cfg.Username, cfg.PG.Password, cfg.PG.Host, cfg.PG.Port, cfg.PG.Name)
//		m, err := migrate.New("file://migrations", url)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		if err := m.Up(); err != nil && err.Error() != "no change" {
//			log.Fatal(err)
//		}
//	}
//
//	func init() {
//		cfg, err := config.New()
//		if err != nil {
//			log.Fatalf("Config error: %s", err)
//		}
//		migrateUp(cfg)
//
//		fmt.Println("migrated")
//
// }

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host       localhost:8000
// @BasePath  /v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
