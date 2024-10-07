dep:
	go mod download
	go mod tidy

wire:
	wire ./...
