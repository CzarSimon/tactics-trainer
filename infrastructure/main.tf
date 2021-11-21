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
    "slow_query_log"  = "ON"
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

resource "scaleway_instance_security_group" "bastion_host_sg" {
  name                    = "bastion-host-security-group"
  inbound_default_policy  = "drop"
  outbound_default_policy = "drop"

  inbound_rule {
    action = "accept"
    port   = 22
    ip     = "188.151.131.138"
  }

  inbound_rule {
    action = "accept"
    port   = 22
    ip     = "185.201.174.65"
  }

  outbound_rule {
    action = "accept"
    port   = scaleway_rdb_instance.main.endpoint_port
    ip     = scaleway_rdb_instance.main.endpoint_ip
  }
}

resource "scaleway_instance_ip" "bastion_host_ip" {}

resource "scaleway_rdb_acl" "allow_bastion_host" {
  instance_id = scaleway_rdb_instance.main.id
  acl_rules {
    ip          = join("/", [scaleway_instance_ip.bastion_host_ip.address, "32"])
    description = "bastion host ip"
  }
}

resource "scaleway_instance_server" "bastion_host" {
  name              = "bastion-host"
  type              = "DEV1-S"
  image             = "debian_buster"
  ip_id             = scaleway_instance_ip.bastion_host_ip.id
  security_group_id = scaleway_instance_security_group.bastion_host_sg.id
  tags              = ["tactics-trainer", "bastion-host"]
  user_data = {
    cloud-init = file("${path.module}/cloud-init.yml")
  }
}
