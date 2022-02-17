go-run:
	@sh ./scripts/go-run.sh

swag-init:
	swag init --parseInternal --exclude build,developments,docs,scripts,vendor -g cmd/efishery-golang/main.go 

swag-fmt:
	swag fmt --exclude build,developments,docs,scripts -g cmd/efishery-golang/main.go