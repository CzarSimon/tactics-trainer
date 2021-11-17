test:
	go vet main.go
	gosec .
	trivy fs --ignore-unfixed --exit-code 1 .
	trivy config --exit-code 1 .


run-local:
	SCALEWAY_REGION='fr-par' ACL_IGNORE='bastion host ip' go run main.go