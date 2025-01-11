build:
	@echo "Initializing demo git repository..."
	git init
	git add .
	git commit -m "Initial commit"
	@echo "Building the project..."
	go build -o main
	@echo "Running the project..."
	./main
	@echo "Cleaning up..."
	rm -rf build-info
	rm -rf .git
	@echo "Done!"