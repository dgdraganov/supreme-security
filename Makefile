

compose:
	docker-compose up --build

decompose:
	docker-compose down

server:
	docker-compose up --build server

client:
	docker-compose up --build client
