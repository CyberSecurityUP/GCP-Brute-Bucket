# Binary name
BINARY_NAME = gcp-brute-bucket

# Project dependencies
DEPS = github.com/fatih/color

all: build

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go get -u $(DEPS)

# Build the binary
build: deps
	@echo "Building the binary..."
	@go build -o $(BINARY_NAME) main.go

# Run the program with an example wordlist
run: build
	@echo "Running the program..."
	@./$(BINARY_NAME) wordlist.txt

# Clean generated files
clean:
	@echo "Removing generated files..."
	@rm -f $(BINARY_NAME)

# Update Go modules
update:
	@echo "Updating Go modules..."
	@go mod tidy
