build:
	mkdir -p bin
	go build -o bin ./cmd/main.go

run:
	go get github.com/go-sql-driver/mysql
	go run ./cmd/main.go

compile:
	mkdir -p bin
	GOOS=freebsd GOARCH=386 go build -o bin/bsd ./cmd/main.go
	GOOS=linux GOARCH=386 go build -o bin/linux ./cmd/main.go
	GOOS=windows GOARCH=386 go build -o bin/windows ./cmd/main.go



