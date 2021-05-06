run:
	go run main.go

docker:
	docker build -t kindergarten .
	docker run --name=myapp --rm kindergarten
	docker rmi kindergarten