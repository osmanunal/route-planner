devdb:
	docker-compose -f docker-compose.yml up mysql -d

migrate:
	go run pkg/cmd/main.go migrate

resetdb:
	go run pkg/cmd/main.go resetdb
