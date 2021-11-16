terraform {
  required_providers {
    scaleway = {
      source = "scaleway/scaleway"
    }
  }
  required_version = ">= 0.13"
}

provider "scaleway" {
  zone       = "fr-par-1"
  region     = "fr-par"
  project_id = "0bc19c59-3a7b-44bb-b7a2-53d44c07cf9c"
}

resource "null_resource" "kubeconfig" {
  depends_on = [
    scaleway_k8s_pool.main
  ]
  triggers = {
    host                   = scaleway_k8s_cluster.main.kubeconfig[0].host
    token                  = scaleway_k8s_cluster.main.kubeconfig[0].token
    cluster_ca_certificate = scaleway_k8s_cluster.main.kubeconfig[0].cluster_ca_certificate
  }
}

provider "kubernetes" {
  host  = null_resource.kubeconfig.triggers.host
  token = null_resource.kubeconfig.triggers.token
  cluster_ca_certificate = base64decode(
    null_resource.kubeconfig.triggers.cluster_ca_certificate
  )
}

resource "kubernetes_namespace" "application" {
  depends_on = [
    scaleway_k8s_pool.main
  ]
  metadata {
    name = "application"
  }
}

resource "kubernetes_namespace" "monitoring" {
  depends_on = [
    scaleway_k8s_pool.main
  ]
  metadata {
    name = "monitoring"
  }
}

resource "kubernetes_namespace" "cert-manager" {
  depends_on = [
    scaleway_k8s_pool.main
  ]
  metadata {
    name = "cert-manager"
  }
}

resource "kubernetes_namespace" "argo" {
  depends_on = [
    scaleway_k8s_pool.main
  ]
  metadata {
    name = "argo"
  }
}

resource "scaleway_rdb_instance" "main" {
  name              = "tactics-trainer-db"
  node_type         = "db-dev-s"
  engine            = "MySQL-8"
  is_ha_cluster     = false
  disable_backup    = false
  volume_type       = "bssd"
  volume_size_in_gb = 10
  tags              = ["tactics-trainer"]
  settings = {
    "max_connections" = "100"
  }
}

resource "scaleway_rdb_database" "puzzle-server-db" {
  instance_id = scaleway_rdb_instance.main.id
  name        = "puzzleserver"
}

resource "scaleway_rdb_database" "iam-server-db" {
  instance_id = scaleway_rdb_instance.main.id
  name        = "iamserver"
}

resource "scaleway_rdb_acl" "main" {
  instance_id = scaleway_rdb_instance.main.id
  acl_rules {
    ip          = "188.151.131.138/32"
    description = "home ip"
  }
}

resource "scaleway_k8s_cluster" "main" {
  name    = "tactics-trainer-cluster"
  version = "1.22"
  cni     = "calico"
  tags    = ["tactics-trainer"]
  auto_upgrade {
    enable                        = true
    maintenance_window_day        = "wednesday"
    maintenance_window_start_hour = "4"
  }
}

resource "scaleway_k8s_pool" "main" {
  cluster_id  = scaleway_k8s_cluster.main.id
  name        = "main-pool"
  node_type   = "DEV1-M"
  size        = 1
  min_size    = 1
  max_size    = 2
  autoscaling = true
  autohealing = true
}
