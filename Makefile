#!bin/bash
run-admin:
	go run cmd/admin/main.go

swagger-admin:
	swag init -d ./ -g cmd/admin/main.go \
    -o ./docs/admin --pd

# delete submodules folder in git cache
# git rm --cached submodules