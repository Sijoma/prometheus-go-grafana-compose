run:	
	docker-compose up --build 

daemon:
	docker-compose up -d --build

stop-daemon:
	docker-compose down

reset:
	docker-compose down -v

fire:
	hey -n 9000 http://localhost:9200