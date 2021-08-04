.PHONY: generate-migration build

build:
	go build -o todo-project

generate-migration:
	docker run \
	-v "${PWD}"/migrations:/migrations:rw \
  --network host \
  --user 1000:1000 \
  migrate/migrate \
  create \
  --ext sql \
  --dir /migrations/ \
  todo-app