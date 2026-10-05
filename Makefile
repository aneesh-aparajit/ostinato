ORANGE := \033[38;5;208m
GREEN  := \033[32m
RESET  := \033[0m

.PHONY: build run

build:
	@printf "$(ORANGE)building ostinato...$(RESET)\n"
	@go build -o bin/ostinato

run: build
	@printf "$(GREEN)running ostinato...$(RESET)\n"
	@./bin/ostinato
