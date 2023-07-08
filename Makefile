# dev:
# 	@docker ps | grep mysql_test_app > /dev/null || make db_mysql_dev && \
# 	@docker ps | grep redis_test_app > /dev/null || make redis_dev && \
# 	air

dev:
	@docker ps | grep postgres_test_app > /dev/null || make db_postgres_dev && \
	air

test:
	@go test -v -count=1 ./...

db_mysql_dev:
	@echo "Starting mysql container..." && \
	docker run --rm --name mysql_test_app \
		-v $(PWD)/.docker/mysql:/var/lib/mysql \
		-e MYSQL_ROOT_PASSWORD=secret \
		-e MYSQL_DATABASE=test_app \
		-p 3306:3306 \
		-d mysql && \
	echo "Waiting for mysql to start..." && \
	sleep 5

db_postgres_dev:
	@echo "Starting postgres container..." && \
	docker run --rm --name postgres_test_app \
		-v $(PWD)/.docker/postgres:/var/lib/postgresql/data \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=test_app \
		-p 5432:5432 \
		-d postgres && \
	echo "Waiting for postgres to start..." && \
	sleep 5

redis_dev:
	@echo "Starting redis container..." && \
	docker run --rm --name redis_test_app \
		-v $(PWD)/.docker/redis:/data \
		-p 6379:6379 \
		-d redis

stop:
	@echo "Stopping containers..." && \
	docker stop mysql_test_app redis_test_app postgres_test_app

clean_db:
	@make stop && \
	echo "Cleaning up database files..." && \
	@rm -rf ./.docker/

.PHONY: dev test db_mysql_dev db_postgres_dev redis_dev stop clean_db
