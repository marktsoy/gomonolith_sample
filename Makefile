.PHONY: build
build:
	go build -v ./cmd/app/main.go


.DEFAULT_GOAL:=build
 

.PHONY: migrate
migrate:
	migrate -database postgres://localhost/monolith_tg_sender?sslmode=disable -path migrations up
	migrate -database postgres://localhost/sandbox?sslmode=disable -path migrations up
	


.PHONY: resetdb
resetdb:
	migrate -database postgres://localhost/monolith_tg_sender?sslmode=disable -path migrations down
	migrate -database postgres://localhost/sandbox?sslmode=disable -path migrations down

.PHONY: test
test:
	go test -v -race -timeout 30s ./...