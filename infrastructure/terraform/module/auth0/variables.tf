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

variable "callbacks" {
  type = list(string)
}

variable "allowed_logout_urls" {
  type = list(string)
}

variable "web_origins" {
  type = list(string)
}
