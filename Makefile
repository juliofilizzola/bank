up:
	migrate -path ./internal/db/migration -database mysql://banker:123456@tcp/bank-dev --verbose up

up-test:
	migrate -path ./internal/db/migration -database mysql://root:123456@tcp/bank-test --verbose up

test-file:
	go tool cover --func=cover.txt

view-test:
	go tool cover --html=cover.txt

test-cover:
	go test ./... -coverprofile cover.txt