up:
	docker exec -i ecomn-mysql mysql -uroot -ppassword <<< "CREATE DATABASE ecomm";
	docker run -it --rm --network host --volume ./db:/db migrate/migrate -path=/db/migrations -database "mysql://root:password@tcp(localhost:3306)/ecomm" up
down:
	docker run -it --rm --network host --volume ./db:/db migrate/migrate -path=/db/migrations -database "mysql://root:password@tcp(localhost:3306)/ecomm" down
