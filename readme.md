# BloGo

![blogo](https://github.com/user-attachments/assets/73108bb4-5b7a-4255-ac74-d99ca8da3c1e)

Requirements: goose(schema), sqlc(queries), docker, postgresql

Db cmds

``` bash
docker pull postgres

docker run --name postgresdb   -e POSTGRES_USER=myuser   -e POSTGRES_PASSWORD=mysecretpassword   -e POSTGRES_DB=postgresdb   -v $(pwd)/db:/docker-entrypoint-initdb.d/   -p 5432:5432   -d postgres

docker run -p 80:80 \
  -e "PGADMIN_DEFAULT_EMAIL=user@domain.com" \
  -e "PGADMIN_DEFAULT_PASSWORD=admin" \
  -d dpage/pgadmin4


goose postgres "host=localhost port=5432 user=myuser password=mysecretpassword dbname=blogoDb sslmode=disable" up

sqlc generate # scafolds database & models
```

Request auth requires `Header: Authorization : ApiKey 140194325234053`
