phony: run
run:
	docker-compose up my-client

phony: run-server
run-server:
	docker-compose up my-server


phony: clean
clean:
	docker-compose down my-server my-client
	@bash -c 'h=$$(docker images "*-server" -q); if [ -n "$$h" ]; then docker rmi $$h; fi'
	@bash -c 'h=$$(docker images "*-client*" -q); if [ -n "$$h" ]; then docker rmi $$h; fi'

phony: test
test:
	go test -v ./...