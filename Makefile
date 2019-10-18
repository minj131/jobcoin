# LOGGING
LOG_MSG=$(shell command -v figlet 2> /dev/null)
ifeq ($(LOG_MSG),)
LOG_MSG=echo
endif

# Color printing of output
PRINT_BLUE=awk '{ print "\033[1;34m" $$0 "\033[1;0m" }'

build: ## Builds binary
	@$(LOG_MSG) Building binary ... | $(PRINT_BLUE)
	mkdir -p bin && go build -o bin/jobcoin -i github.com/minj131/jobcoin

unit-tests: ## Runs unit tests
	@$(LOG_MSG) Running unit tests ... | $(PRINT_BLUE)
	@scripts/unit_test