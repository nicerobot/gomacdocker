.PHONY : $(DIRS)
.PHONY : app run test
.PHONY : help
.DEFAULT_GOAL := help


app: main.go ## Make everything
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@
	docker build --tag=test .

run: ## Run the container
	docker run -it -p 8008:8008 test debug

test: ## Run test
	go test -v


help: ## This help.
	@echo Targets:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_ -]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
