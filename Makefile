clean:
	@go clean -testcache

format:
	gofmt -w -l .

lint:
	@$(MAKE) format
	@./scripts/lint_go_mods.sh

cover:
	@go test -covermode=count -coverprofile=geral.out.tmp ./... && cat geral.out.tmp | grep -v fake > geral.out && go tool cover -func=geral.out

test:
	@$(MAKE) -s clean
	@go test -race $(shell go list ./... | grep -Ev 'script|cmd')
	@$(MAKE) -s cover

build-linux:
	GOOS=linux go build -o myapi ./cmd/myapi/main.go