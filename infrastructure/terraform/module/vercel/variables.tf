variable "vercel_api_token" {
  type      = string
  sensitive = true
}

variable "auth0_domain" {
  type = string
}

variable "auth0_client_id" {
  type = string
}
