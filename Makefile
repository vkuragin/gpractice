all: gpractice gpractice-web

gpractice: ./cmd/gpractice/main.go
	go install -v ./cmd/gpractice/

gpractice-web: ./cmd/gpractice-web/main.go
	go install -v ./cmd/gpractice-web/