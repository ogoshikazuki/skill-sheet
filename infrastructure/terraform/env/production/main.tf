module "auth0" {
  source = "../../module/auth0"

  auth0_domain         = var.auth0_domain
  auth0_client_id      = var.auth0_client_id
  auth0_client_secret  = var.auth0_client_secret
  auth0_admin_email    = var.auth0_admin_email
  auth0_admin_password = var.auth0_admin_password
  callbacks            = ["https://${module.vercel.admin_domain}"]
  allowed_logout_urls  = ["https://${module.vercel.admin_domain}"]
  web_origins          = ["https://${module.vercel.admin_domain}"]
}

module "vercel" {
  source = "../../module/vercel"

  vercel_api_token = var.vercel_api_token
  auth0_domain     = var.auth0_domain
  auth0_client_id  = module.auth0.client_id
}

module "railway" {
  source = "../../module/railway"

  token = var.railway_token
}
