start:
	minikube start  &&\
	minikube addons enable metrics-server &&\
    helm dependencies build deployments/k8s/helm/kuberrest &&\
	helm install kuberrest deployments/k8s/helm/kuberrest &&\
	minikube dashboard
.PHONY: start
##kubectl port-forward svc/kuberrest-grafana 8080:3000
done:
	helm delete kuberrest && \
	kubectl delete jobs kuberrest-job && \
	minikube stop
.PHONY: done

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


#helm upgrade --install ingress-nginx ingress-nginx \
#  --repo https://kubernetes.github.io/ingress-nginx \
#  --namespace ingress-nginx --create-namespace


#helm upgrade --install ingress-nginx ingress-nginx \
#--repo https://kubernetes.github.io/ingress-nginx\
#--namespace ingress-nginx --create-namespace
#--set controller.metrics.enabled=true \
#--set controller.metrics.serviceMonitor.enabled=true \
#--set controller.metrics.serviceMonitor.additionalLabels.release="prometheus"

#helm upgrade prometheus prometheus-community/kube-prometheus-stack \
#--set prometheus.prometheusSpec.podMonitorSelectorNilUsesHelmValues=false \
#--set prometheus.prometheusSpec.serviceMonitorSelectorNilUsesHelmValues=false

