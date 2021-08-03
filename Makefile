dev-serve:
	go run ./cmd/server/main.go -port 9091

dev-client:
	go run ./cmd/client/main.go -address 0.0.0.0:9091


dev-pis:
	go run main.go
	
genfiles:
	protoc --proto_path=internal/proto internal/proto/*.proto  --go_out=:internal/proto --go-grpc_out=:internal/proto --grpc-gateway_out=:internal/proto --openapiv2_out=:pkg/swagger

migrateup:
	migrate -path internal/db/migration -database "postgresql://postgres:postgres@tgpisdevdb.chqiulfy2dsu.us-east-1.rds.amazonaws.com:5432/tg_pis_db?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://postgres:postgres@tgpisdevdb.chqiulfy2dsu.us-east-1.rds.amazonaws.com:5432/tg_pis_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

clean:
	rm internal/proto/integration/*
	rm pkg/swagger/*

.PHONY:
	postgres createdb dropdb migrateup migratedown sqlc test server