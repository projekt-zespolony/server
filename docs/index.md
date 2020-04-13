# Getting started

Start the database:

```sh
docker run -d --rm --name db -e MYSQL_DATABASE=db -e MYSQL_USER=user -e MYSQL_PASSWORD=pass -e MYSQL_RANDOM_ROOT_PASSWORD=yes mysql:8
```

Build the server image:

```sh
docker build -t server:latest .
```

Start the server:

```sh
docker run -d --rm --name server --links db -p 8080:80 -e DB_NAME=db -e DB_USER=user -e DB_PASS=pass -e DB_ADDR=db -e SERVER_PORT=80 -e SERVER_ACCESS_TOKEN=token server:latest
```
