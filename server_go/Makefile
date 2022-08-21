debug:
	GOOS=linux GOARCH=amd64 \
		go build -o bin/. cmd/server_go/main.go

run:
	GOOS=linux GOARCH=amd64 \
		go run cmd/server_go/main.go

release:
	GIN_MODE=release GOOS=linux GOARCH=amd64 \
		go build -o bin cmd/server_go/main.go