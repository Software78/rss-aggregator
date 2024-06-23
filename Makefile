
help: 
	@echo "Usage: make [command]"
	@echo ""
	@echo "Commands:"
	@echo "  build     Build the application"

build:
	@echo "Building..."
	@go build && ./rss-aggregator
	@echo "Build complete"