up:
	migrate -path db/migration -database postgres://banker:123456@localhost:5432/bank-dev?sslmode=disable --verbose up