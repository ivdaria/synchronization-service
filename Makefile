DB_DSN="postgres://postgres:postgres@localhost:5432/algorithms"
MIGRATIONS_DIR=migrations

install-deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go install go.uber.org/mock/mockgen@latest

migrations-up:
	goose -dir ${MIGRATIONS_DIR} postgres ${DB_DSN} up

migrations-down:
	goose -dir ${MIGRATIONS_DIR} postgres ${DB_DSN} down

migrations-status:
	goose -dir ${MIGRATIONS_DIR} postgres ${DB_DSN} status
