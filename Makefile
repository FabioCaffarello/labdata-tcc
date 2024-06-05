

build-docs:
	npx nx graph --file=docs/dependency-graph/index.html
	npx nx  run-many --target=godoc --all

serve-doc: build-docs
	poetry run mkdocs serve

deploy-doc: build-docs
	poetry run mkdocs gh-deploy

lint:
	npx nx run-many --target=lint --all

check-all:
	npx nx run-many --target=test --all
