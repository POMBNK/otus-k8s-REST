generate:
	go generate ./...
	go run doc/merger/main.go
	statik -src=doc/dist -include=*.html,*.css,*.js,*.png,*.json
.PHONY: generate

mup:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres up
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres status
.PHONY: mup

mdown:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres down
	goose -dir migrations postgres postgres://postgres:postgres@localhost:5432/postgres status
.PHONY: mdown