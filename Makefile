GOOS=darwin
GOARCH=arm64
GOENGINE=go

debug:
	GOOS=${GOOS} GOARCH=${GOARCH} ${GOENGINE} build -o bin/. cmd/server_go/main.go 
	./bin/main

run:
	GOOS${GOOS} GOARCH=${GOARCH} ${GOENGINE} run cmd/server_go/main.go

release:
	GOOS=${GOOS} GOARCH=${GOARCH} ${GOENGINE} build -o bin/. cmd/server_go/main.go
	GIN_MODE=release ./bin/main