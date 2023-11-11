migrate-docker:
	migrate -path ./migrations -database "postgres://fr:fr@db:6543/fr?sslmode=disable" up
