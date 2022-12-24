.PHONY: run migrate_up migrate_down

# Go migrate
migrate_up:
	migrate -path migrations -database "mysql://root:109339Lam@@tcp(127.0.0.1:3307)/go-airbnb" up
migrate_down:
	migrate -path migrations -database "mysql://root:109339Lam@@tcp(127.0.0.1:3307)/go-airbnb" down
run:
	go run main.go	