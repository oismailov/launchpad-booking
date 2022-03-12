.PHONY: start stop init build tests migrate refresh-upcoming-launches refresh-launchpad-list refresh

start:
	docker-compose up -d

stop:
	docker-compose stop

init:
	echo "to be defined"

build:
	echo "to be defined"

tests:
	echo "to be defined"

migrate:
	docker-compose exec app go run migration/migration.go

refresh-upcoming-launches:
	docker-compose exec app go run cmd/launches/bin/run.go

refresh-launchpad-list:
	docker-compose exec app go run cmd/launchpads/bin/run.go

refresh:
	docker-compose exec app go run cmd/launches/bin/run.go
	docker-compose exec app go run cmd/launchpads/bin/run.go

ssh:
	docker-compose exec app bash

