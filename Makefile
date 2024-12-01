migration_up:
	migrate -source "file://$(shell pwd)/db/migrations" -database "mysql://miniblog_dev:dev@tcp(127.0.0.1)/miniblog" up
