## Include variables from the .envrc file
include .envrc

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@export DATABASE_URL=${DATABASE_URL} && go run ./cmd/ 