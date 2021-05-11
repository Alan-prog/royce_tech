build:
	docker container rm --force royce_application 2>/dev/null && docker build -t royce_application . && docker run --name royce_application -e POSTGRES_PASSWORD=somepass -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres --rm -p 8080:8080 -d royce_application
run:
	docker exec -it royce_application /royce

