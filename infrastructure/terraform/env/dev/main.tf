locals {
  admin_origin = "http://localhost:3000"
}

module "auth0" {
  source = "../../module/auth0"

  auth0_domain         = var.auth0_domain
  auth0_client_id      = var.auth0_client_id
  auth0_client_secret  = var.auth0_client_secret
  auth0_admin_email    = var.auth0_admin_email
  auth0_admin_password = var.auth0_admin_password
  callbacks            = [local.admin_origin]
  allowed_logout_urls  = [local.admin_origin]
  web_origins          = [local.admin_origin]
}
