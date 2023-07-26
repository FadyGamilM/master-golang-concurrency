# for migration commands
migrate_up:
	docker run -i -v "/h/1- freelancing path/Courses/golang stack/projects/master-golang-concurrency/final-project-subscribtion-service/db/migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://concurrency:concurrency@127.0.0.1:5432/concurrencyDB?sslmode=disable" up 1

migrate_down:
	docker run -i -v "/h/1- freelancing path/Courses/golang stack/projects/master-golang-concurrency/final-project-subscribtion-service/db/migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "postgresql://concurrency:concurrency@127.0.0.1:5432/concurrencyDB?sslmode=disable" down 1
