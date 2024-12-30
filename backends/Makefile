## Include variables from the .envrc file
include .envrc

## run/api: run the cmd/api application
.PHONY: run/api/ride
run/api/ride:
	@cd ride && go run ./... 

.PHONY: run/api/account
run/api/account:
	@cd ./account && go run ./...


.PHONY: run/api/payment
run/api/payment:
	@cd ./payment && go run ./...