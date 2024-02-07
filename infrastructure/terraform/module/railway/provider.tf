terraform {
  required_providers {
    railway = {
      source  = "terraform-community-providers/railway"
      version = "0.3.1"
    }
  }
}

provider "railway" {
  token = var.token
}
