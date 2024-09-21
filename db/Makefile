migrate:
	docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate create -ext sql -dir /db/migrations init_schema
session_table:
	docker run -it --rm --network host --volume "$(pwd)/db:/db" migrate/migrate create -ext sql -dir /db/migrations add_session_table
