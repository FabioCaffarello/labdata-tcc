WORKSPACE=labdata-tcc


guard-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Environment variable $* not set"; \
		exit 1; \
	fi

build-docs:
	npx nx graph --file=docs/dependency-graph/index.html
	npx nx  run-many --target=godoc --all

serve-doc: build-docs
	poetry run mkdocs serve

deploy-doc: build-docs
	poetry run mkdocs gh-deploy

lint:
	npx nx run-many --target=lint --all

check: guard-project cleanup
	npx nx test $(project)

check-all: cleanup
	npx nx run-many --target=test --all

go-image:
	npx nx run-many --target=image --projects=tag:lang:golang

py-image:
	npx nx run-many --target=build --projects=tag:lang:python --devDependencies
	npx nx run-many --target=image --projects=tag:lang:python

image: go-image py-image
	echo "Images built"

chech-integration-all:
	npx nx run-many --target=check-integration --all

chech-integration: guard-project
	npx nx check-integration $(project)

run:
	docker-compose up -d

purge-images:
	@docker images --filter "dangling=true" -q | xargs -r docker rmi

install:
	npx nx run-many --target=install --with dev --all


cleanup:
	@npx nx reset;
	@containers=$$(docker ps -q -a); \
	if [ -n "$$containers" ]; then \
		docker rm -f $$containers; \
	else \
		echo "No containers to remove"; \
	fi

PHONY: build-docs serve-doc deploy-doc lint check check-all cleanup