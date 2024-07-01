Curso: https://www.udemy.com/course/golang-postgres-e-react-num-projeto-fullstack-de-financas/

Comando para criar o container com mysql
`docker run -d --name postgres-go-next -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=finance -e MYSQL_USER=golang -e MYSQL_PASSWORD=golang mysql:latest`
Comando postgres
`docker run --name postgres14 -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine`

criar a tabela de migrations
`migrate create -ext sql -dir db/migration -seq initial_tables`

`docker  exec -it postgres /bin/sh`
`createdb --username=postgres --owner=postgres go_finance`

`migrate -path db/migration -database "postgresql://postgres:postgres@172.17.0.3:5432/go_finance?sslmode=disable" -verbose up`
