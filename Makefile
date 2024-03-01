GOOS=darwin
GOARCH=arm64

debug:
	${GOOS} ${GOARCH} go build -o bin/. cmd/server_go/main.go 
	./bin/main

run:
	${GOOS} ${GOARCH} go run cmd/server_go/main.go

release:
	${GOOS} ${GOARCH} go build -o bin/. cmd/server_go/main.go
	GIN_MODE=release ./bin/main