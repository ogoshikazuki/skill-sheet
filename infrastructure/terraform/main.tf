module "vercel" {
  source = "./module/vercel"

  vercel_api_token = var.vercel_api_token
}

module "railway" {
  source = "./module/railway"

  token = var.railway_token
}
