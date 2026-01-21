.PHONY: help run build

help:
	@echo "Available commands:"
	@echo "  run    : Run the web application with live-reloading"
	@echo "  build  : Compile the web application"

run:
	@echo "Starting web application with air..."
	@RESEND_API_KEY=re_NZvsQjJ2_HQws5BB5hUpachdkcJtxGv75 CONTACT_EMAIL=nelsonmarro99@gmail.com air

build:
	@echo "Generating templ files and compiling Go application..."
	@templ generate
	@go build -o ./bin/server ./cmd/web/main.go
