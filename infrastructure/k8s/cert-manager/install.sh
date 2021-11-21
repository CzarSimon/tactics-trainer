helm install cert-manager \
  jetstack/cert-manager \
  --namespace cert-manager \
  --version v1.6.0 \
  --set installCRDs=true
