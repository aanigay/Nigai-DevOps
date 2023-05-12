.PHONY: run
run:
	go run cmd/service/main.go

.PHONY: dc-up
dc-up:
	docker-compose up -d --remove-orphans --build

.PHONY: dc-stop
dc-stop:
	docker-compose stop

.PHONY: pg-up
pg-up:
	docker-compose -f docker-compose.db.yml up -d --remove-orphans

.PHONY: pg-stop
pg-stop:
	docker-compose -f docker-compose.db.yml stop