.PHONY: help run build

help:
	@echo "Available commands:"
	@echo "  run    : Run the web application with live-reloading"
	@echo "  build  : Compile the web application"

run:
	@echo "Starting web application with air..."
	@air

build:
	@echo "Generando assets y compilando para Linux..."
	npm run build:css
	npm run build:js
	templ generate
	GOOS=linux GOARCH=amd64 go build -o bin/server ./cmd/web/main.go
