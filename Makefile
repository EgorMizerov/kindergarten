run:
	go run main.go

docker:
	docker build -t kindergarten .
	docker run --name=myapp --rm -p 8000:8000 kindergarten
	docker rmi kindergarten