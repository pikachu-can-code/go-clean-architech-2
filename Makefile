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

init_test:
	docker-compose exec mysql /etc/mysql/test/create_db_test.sh
	docker-compose run --rm --entrypoint "./scripts/migrate_test.sh up" app

gen_proto:
	docker-compose run --rm --entrypoint "buf generate" app-rpc-gen

tidy:
	docker-compose run --rm --entrypoint "go mod tidy" app
