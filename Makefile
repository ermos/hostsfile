test:
	@go test

test/coverage:
	@go test -cover

test/html:
	@go test -coverprofile=coverage.out
	@go tool cover -html=coverage.out