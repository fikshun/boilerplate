build:
	docker buildx build . -t <registry>/go-backend:$$(cat version.txt) --platform linux/amd64 --load
	docker tag <registry>/go-backend:$$(cat version.txt) ghcr.io/<registry>/go-backend:latest
	docker push ghcr.io/<registry>/go-backend:$$(cat version.txt)
	docker push ghcr.io/<registry>/go-backend:latest
build-snapshot:
	docker build . -t ghcr.io/<registry>/go-backend:SNAPSHOT-$$(git rev-parse --short HEAD)-$$(cat version.txt)
build-local:
	docker build . -t ghcr.io/<registry>/go-backend:local
build-local-restart:
	docker build . -t ghcr.io/<registry>/go-backend:local
	docker-compose down; docker-compose up -d
dev-up:
	docker-compose up -d 
dev-down:
	docker-compose down
dev-restart:
	docker-compose down; docker-compose up -d
# generate boilerplate from openapi - requires https://github.com/deepmap/oapi-codegen
generate:
	oapi-codegen --config=types.cfg.yml openapi.yml
	oapi-codegen --config=server.cfg.yml openapi.yml

