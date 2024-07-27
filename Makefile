start:
	minikube start
.PHONY: start

apply:
	cd deployments/k8s/manifests && kubectl apply -f.
.PHONY: apply

tun:
	minikube tunnel
.PHONY: tunnel

generate:
	go generate ./...
	go run doc/merger/main.go
.PHONY: generate

mup:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:6969/postgres up
	goose -dir migrations postgres postgres://postgres:postgres@localhost:6969/postgres status
.PHONY: mup

mdown:
	goose -dir migrations postgres postgres://postgres:postgres@localhost:6969/postgres down
	goose -dir migrations postgres postgres://postgres:postgres@localhost:6969/postgres status
.PHONY: mdown