.PHONY: compose-up compose-down

compose-up:
	docker compose -f deploy/compose/docker-compose.dev.yml up

compose-down:
	docker compose -f deploy/compose/docker-compose.dev.yml down
