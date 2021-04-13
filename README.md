# go-todo-api

### set up
```
// start up
$docker-compose build
$docker-compose up -d

// go inside container
$docker exec -it go-todo /bin/sh         // go-todo
$docker exec -it mysql-container /bin/sh // mysql

// run
$docker-compose exec app go run main.go

// open mysql
// fron inside mysql-container
$mysql -utodo -p
```
