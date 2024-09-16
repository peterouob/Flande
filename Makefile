up:
	docker exec -i ecomn-mysql mysql -uroot -ppassword <<< "CREATE DATABASE ecomm";
	docker run -it --rm --network host --volume ./db:/db migrate/migrate -path=/db/migrations -database "mysql://root:password@tcp(localhost:3306)/ecomm" up
down:
	docker run -it --rm --network host --volume ./db:/db migrate/migrate -path=/db/migrations -database "mysql://root:password@tcp(localhost:3306)/ecomm" down
gen_user:
	protoc --go_out=. --go-grpc_out=. rpc/user.proto
kafka_ui:
	docker pull provectuslabs/kafka-ui
