## Include variables from the .envrc file
include .envrc

## run/api: run the cmd/api application
.PHONY: run/api/ride
run/api/ride:
	@export DATABASE_URL=${DATABASE_URL} && go run ./cmd/ride/ 

.PHONY: run/api/account
run/api/account:
	@export DATABASE_URL=${DATABASE_URL} && go run ./cmd/account/


.PHONY: run/api/payment
run/api/payment:
	@export DATABASE_URL=${DATABASE_URL} && go run ./cmd/account/