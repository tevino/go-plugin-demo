all: app p1.so

app: ./cmd/app/main.go
	go build -o $@ $^

p1.so: ./plugins/p1/p1.go
	go build -buildmode plugin -o $@ $^
