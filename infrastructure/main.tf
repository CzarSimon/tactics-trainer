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
