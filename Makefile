include compose/.env
include variables.mk

run:
	DATABASE=$(database) \
	DBHOST=$(localhost) \
	DBPORT=$(dbport) \
    DBUSER=$(dbuser) \
    DBPASS=$(dbpassword) \
    SERVICEHOST=$(localhost) \
    SERVICEPORT=$(port) \
	APP_NAME=$(service) \
    HIVEMQ=$(localhost):$(hiveport) \
    HUSERNAME=$(husername) \
    HPASSWORD=$(hpassword) \
    HCLIENT=$(clientid) \
    HTOPIC=$(htopic) \
	SECRETKEY=$(secretkey) \
	$(GORUN) cmd/main.go

docker-image:
	make -C docker docker-image

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(build) ./cmd/

docker-start:
	make -C compose start

docker-stop:
	make -C compose stop

kick-start: docker-image docker-start

clean:
	@docker rmi -f $(service)
	@docker rmi -f $(database)
	@docker image prune -f
	

.PHONY: run

