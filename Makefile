.PHONY: up
up:
	docker compose up -d --build
	start "Chrome" http://localhost:8080
