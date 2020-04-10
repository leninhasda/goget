build: test
	go build -o ./dist/goget .

test:
	go test -v ./...