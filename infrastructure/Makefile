test:
	trivy config --exit-code 1 .

argo-web:
	kubectl port-forward svc/argocd-server -n argocd 8080:443