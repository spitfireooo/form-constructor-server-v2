up:
	docker compose up --build form-constructor-server auth
down:
	docker compose down -v && docker system prune -a