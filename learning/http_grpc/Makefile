all: http_gateway docker-image

http_gateway:
	CGO_ENABLED=0 go build -o build/http_gateway -a -ldflags '-extldflags "-static"' main.go

docker-image: http_gateway
	docker build -t http_gateway:v1.0.0 ./build

clean:
	rm build/http_gateway
