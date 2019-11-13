TAG=0.1.0
BINARY=datsimple
NAME=teamone/$(BINARY)
IMAGE=$(NAME):$(TAG)
LATEST=$(NAME):latest

build:
	docker build -t $(IMAGE) .
	docker tag $(IMAGE) $(LATEST)

push:
	docker push $(IMAGE)
	docker push $(LATEST)

up:
	make -C ../ up

rm: stop
	make -C ../ rm

stop:
	make -C ../ stop

logs:
	make -C ../ logs

ps:
	make -C ../ ps

compile:
	go build -v -o datsimple && rm datsimple

make_migration:
	goose -dir ./migrations postgres "user=datsimple password=datsimple dbname=datsimple host=datsimple host=127.0.0.1 port=5432 sslmode=disable" up

migrate_up:
	goose -dir ./migrations postgres "user=datsimple password=datsimple dbname=datsimple host=datsimple host=127.0.0.1 port=5432 sslmode=disable" up

migrate_down:
	goose -dir ./migrations postgres "user=datsimple password=datsimple dbname=datsimple host=datsimple host=127.0.0.1 port=5432 sslmode=disable" down

migrate_status:
	goose -dir ./migrations postgres "user=datsimple password=datsimple dbname=datsimple host=datsimple host=127.0.0.1 port=5432 sslmode=disable" status

migrate_reset:
	goose -dir ./migrations postgres "user=datsimple password=datsimple dbname=datsimple host=datsimple host=127.0.0.1 port=5432 sslmode=disable" reset
