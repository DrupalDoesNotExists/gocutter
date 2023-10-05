# Env vars
PROTO_OUT ?= ./
PROTO_SRC ?= ./proto
PROTO_FILES := $(shell find $(PROTO_SRC) -type f -name "*.proto")
GO_OUT ?= ./app
DOCKER_GO_VERSION ?= "latest"

# Tasks

.PHONY: all
## Make the default list of tasks
all: build

.PHONY: build
## Build the Golang files
build: protobuf
	go build -o $(GO_OUT) ./...

.PHONY: protobuf
## Build Protobuf files
protobuf: $(PROTO_FILES)
	-protoc -I=$(PROTO_SRC) --go_out=$(PROTO_OUT) $(PROTO_FILES)

.PHONY: clean
## Clean the Golang compilation results
clean:
	rm -rf $(GO_OUT)

.PHONY: docker-image
## Build the Docker image. DOCKER_GO_VERSION env var is a build arg
docker-image:
	docker build --build-arg "GO_VERSION=$(DOCKER_GO_VERSION)" .

.PHONY: compose-up
## Run the Docker Compose dev containers
compose-up:
	cd deployments && docker-compose build && docker-compose up

.PHONY: compose-down
## Stop and remove the Docker Compose dev containers
compose-down:
	cd deployments && docker-compose down

.PHONY: help
help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)"
	@echo
	@sed -n -e "/^## / { \
		h; \
		s/.*//; \
		:doc" \
		-e "H; \
		n; \
		s/^## //; \
		t doc" \
		-e "s/:.*//; \
		G; \
		s/\\n## /---/; \
		s/\\n/ /g; \
		p; \
	}" ${MAKEFILE_LIST} \
	| LC_ALL='C' sort --ignore-case \
	| awk -F '---' \
		-v ncol=$$(tput cols) \
		-v indent=19 \
		-v col_on="$$(tput setaf 6)" \
		-v col_off="$$(tput sgr0)" \
	'{ \
		printf "%s%*s%s ", col_on, -indent, $$1, col_off; \
		n = split($$2, words, " "); \
		line_length = ncol - indent; \
		for (i = 1; i <= n; i++) { \
			line_length -= length(words[i]) + 1; \
			if (line_length <= 0) { \
				line_length = ncol - indent - length(words[i]) - 1; \
				printf "\n%*s ", -indent, " "; \
			} \
			printf "%s ", words[i]; \
		} \
		printf "\n"; \
	}' \
	| more $(test "$(uname -s)" == "Darwin" && echo '--no-init --raw-control-chars')
