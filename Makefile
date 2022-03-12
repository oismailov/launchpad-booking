.PHONY: start stop init tests migrate refresh-upcoming-launches refresh-launchpad-list refresh

start:
	docker-compose up -d

stop:
	docker-compose stop

init:
	echo "to be defined"

tests:
	docker-compose exec app go test -v ./tests/...

migrate:
	docker-compose exec app go run migration/run.go

refresh-upcoming-launches:
	docker-compose exec app go run cmd/launches/bin/run.go

refresh-launchpad-list:
	docker-compose exec app go run cmd/launchpads/bin/run.go

refresh:
	docker-compose exec app go run cmd/launches/bin/run.go
	docker-compose exec app go run cmd/launchpads/bin/run.go

ssh:
	docker-compose exec app bash

