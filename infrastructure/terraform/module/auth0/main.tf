resource "auth0_client" "auth0_account_management_api_management_client" {
  name = "Auth0 Account Management API Management Client"
}

resource "auth0_client" "skill_sheet" {
  app_type            = "spa"
  name                = "skill-sheet"
  callbacks           = var.callbacks
  allowed_logout_urls = var.allowed_logout_urls
  web_origins         = var.web_origins
  jwt_configuration {
    alg = "RS256"
  }
}

resource "auth0_client_grant" "skill_sheet" {
  client_id = auth0_client.skill_sheet.id
  audience  = "https://skill-sheet"
  scopes    = ["admin"]
}

resource "auth0_connection" "skill_sheet" {
  name     = "Username-Password-Authentication"
  strategy = "auth0"
}

resource "auth0_connection_clients" "skill_sheet" {
  connection_id = auth0_connection.skill_sheet.id
  enabled_clients = [
    auth0_client.auth0_account_management_api_management_client.id,
    auth0_client.skill_sheet.id,
  ]
}

resource "auth0_user" "admin" {
  depends_on = [
    auth0_connection_clients.skill_sheet
  ]

  connection_name = auth0_connection.skill_sheet.name
  email           = var.auth0_admin_email
  email_verified  = true
  password        = var.auth0_admin_password
}
