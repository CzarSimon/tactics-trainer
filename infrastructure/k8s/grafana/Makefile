install:
	helm install grafana grafana/grafana --namespace monitoring

get-admin-password:
	kubectl get secret --namespace monitoring grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo