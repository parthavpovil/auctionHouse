# Auction Application Makefile

CONTRACT_DIR = contract
GO_CMD_DIR = cmd

.PHONY: help run test-status

help:
	@echo "Available Commands:"
	@echo "  help - Show this help"
	@echo "  run  - Run the Go application"
	@echo "  test - Test status endpoint"

run:
	@echo "Starting application..."
	cd $(GO_CMD_DIR) && go run main.go

test:
	@echo "Testing status:"
	@curl -s http://localhost:8080/blockchain/status || echo "App not running"
