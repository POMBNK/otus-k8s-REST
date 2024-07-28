start:
	minikube start  &&\
	minikube addons enable ingress && \
	minikube addons enable metrics-server &&\
    helm dependencies build deployments/k8s/helm/kuberrest &&\
    helm delete kuberrest && \
	helm install kuberrest deployments/k8s/helm/kuberrest &&\
	minikube dashboard
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