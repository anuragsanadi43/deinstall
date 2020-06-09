all:
	go build

run: main.go
	go run main.go

clean: deinstall
	rm deinstall
	