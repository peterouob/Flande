mysql:
	docker run --name ecomn-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:8.4
migrate:
	docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate create -ext sql -dir /db/migrations init_schema
pull:
	docker pull mysql:8.4
