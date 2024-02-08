terraform {
  required_version = "1.7.2"
  required_providers {
    vercel = {
      source  = "vercel/vercel"
      version = "1.0.0"
    }
  }
}

provider "vercel" {
  api_token = var.vercel_api_token
}
