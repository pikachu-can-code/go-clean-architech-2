env:
	cp .env.example .env && cp .env.example .env-ut

start:
	docker-compose up app

stop:
	docker-compose down

clean:
	docker-compose down --rmi all --volumes

migrate_up:
	docker-compose run --rm --entrypoint "./scripts/migrate.sh up" app

migrate_down:
	docker-compose run --rm --entrypoint "./scripts/migrate.sh down" app
