mysql:
	docker run --name ecomn-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:8.4
migrate:
	docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate create -ext sql -dir /db/migrations init_schema
session_table:
	docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate create -ext sql -dir /db/migrations add_session_table
pull:
	docker pull mysql:8.4
