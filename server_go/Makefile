debug:
	GOOS=darwin GOARCH=amd64 go build -o bin/. cmd/server_go/main.go && \
		./bin/main

run:
	GOOS=darwin GOARCH=amd64 go run cmd/server_go/main.go

release:
	GOOS=darwin GOARCH=amd64 go build -o bin/. cmd/server_go/main.go && \
		GIN_MODE=release ./bin/main