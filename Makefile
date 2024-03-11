phony: build-server run-server build-client run-client clean

build-server:
	docker-compose build server

run-server:
	docker-compose up server

build-client:
	docker-compose build client

run-client:
	docker-compose up client

run-client-spam:
	docker-compose up client-spam

run-client-delay:
	docker-compose up client-delay

clean:
	docker-compose down server client client-spam client-delay
	docker-compose rm -f server client client-spam client-delay
	@bash -c 'h=$$(docker images "*-server" -q); if [ -n "$$h" ]; then docker rmi $$h; fi'
	@bash -c 'h=$$(docker images "*-client*" -q); if [ -n "$$h" ]; then docker rmi $$h; fi'
