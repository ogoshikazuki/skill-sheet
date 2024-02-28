variable "auth0_domain" {
  type = string
}

variable "auth0_client_id" {
  type = string
}

variable "auth0_client_secret" {
  type      = string
  sensitive = true
}

variable "auth0_admin_email" {
  type = string
}

variable "auth0_admin_password" {
  type      = string
  sensitive = true
}

variable "vercel_api_token" {
  type      = string
  sensitive = true
}

variable "railway_token" {
  type      = string
  sensitive = true
}
