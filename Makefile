build:  test
	go build -o goget .

test:
	go test -v ./...