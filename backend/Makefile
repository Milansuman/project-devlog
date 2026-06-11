include .env

MIGRATIONS_DIR := ./migrations

new-migration:
	goose -dir $(MIGRATIONS_DIR) create $(NAME) sql

migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DATABASE_URL) up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DATABASE_URL) down
