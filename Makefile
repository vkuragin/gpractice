all: test gpractice gpractice-web

test:
	go test -v .

gpractice: ./cmd/gpractice/main.go
	go install -v ./cmd/gpractice/

gpractice-web: ./cmd/gpractice-web/main.go
	go install -v ./cmd/gpractice-web/