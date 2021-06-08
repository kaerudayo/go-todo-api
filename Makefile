
# mysql
DBUSER := root
DBPASS := todo
DBPORT := 3307
DBNAME := todo

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: mysql
mysql:
	docker-compose exec mysql mysql -u$(DBUSER) -p$(DBPASS) $(DBNAME)
