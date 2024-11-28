.PHONY: test run image container up down prune

up:
	@docker compose up -d

stop:
	@docker compose down

