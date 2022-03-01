# go-clean-architech-1

golang clean architech 1
![golang clean architech 1](docs/img/docs.jpg)

### Create environment file

`cp .env.example.yml .env.yml`

- Then edit data in .env.yml file for your environment

### Generate protobuf

- Install buf: https://docs.buf.build/installation/
- Install protoc-gen plugins:

```
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

- Generate protobuf:

```
buf generate
```

### DB MIGRATIONS note

install goose

```
go get -u github.com/pressly/goose/cmd/goose
```

go inside migration folder
`cd database/migrations`
to create migration file data value in golang (example: role value), then implement code to new generated file

```
goose create role-value go
```

back to database folder
`cd database`
build our custom goose
`go build -o goose *.go`
run your goose command
`./goose -dir "migrations" "DBSTRING" COMMAND`
command up to update version in db, command down to downgrate version in db

read more from: https://github.com/pressly/goose

- Create migration to insert permission to DB:

```bash
cd database/migrations
goose create create-<permission-name> go
```

- Add code to insert permission:

```go
// Up
tx.Exec("INSERT INTO permissions (id, name, created_at, updated_at) VALUES (1, 'ping', NOW(), NOW());")
// Down
tx.Exec("DELETE from permissions WHERE name='ping';")
```
