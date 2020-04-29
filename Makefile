.PHONY: test tidy

test:
	go test --cover ./...

tidy:
	go mod tidy
