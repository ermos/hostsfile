lint:
	@go fmt . && golangci-lint run

test:
	@go test

test/coverage:
	@go test -cover

test/html:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out

generate/examples:
	@cd ./examples/add-host && go run .
	@cd ./examples/edit-host && go run .
	@cd ./examples/remove-host && go run .
	@cd ./examples/add-raw && go run .
	@cd ./examples/clear-hosts-file && go run .
	@cd ./examples/new-hosts-file && go run .