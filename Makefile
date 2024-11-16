.PHONY: build

build:
	go mod tidy && \
	templ generate && \
	tailwind -o public/static/main.css -i public/preset.css && \
	go build -o build/server cmd/server.go