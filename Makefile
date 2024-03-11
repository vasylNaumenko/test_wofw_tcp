phony: build-server run-server build-client run-client clean

build-server:
	docker-compose build server

run-server:
	docker-compose up server

build-client:
	docker-compose build client

run-client:
	docker-compose up client

clean:
	docker-compose down server client
	docker-compose rm -f server client
	@bash -c 'h=$$(docker images "*-server" -q); if [ -n "$$h" ]; then docker rmi $$h; fi'
	@bash -c 'h=$$(docker images "*-client" -q); if [ -n "$$h" ]; then docker rmi $$h; fi'
