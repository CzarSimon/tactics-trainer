helm install cert-manager \
  jetstack/cert-manager \
  --namespace cert-manager \
  --version v0.15.1 \
  --set installCRDs=true
