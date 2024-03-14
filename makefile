
.PHONY: migration-generate
migration-generate:
	docker compose run --rm migrate create -ext sql -dir ./migrations -seq $(name) 
