locals {
  git_repository = {
    type              = "github"
    repo              = "ogoshikazuki/skill-sheet"
    production_branch = "main"
  }
}

resource "vercel_project" "skill_sheet" {
  name           = "skill-sheet"
  framework      = "nuxtjs"
  git_repository = local.git_repository
  ignore_command = "if [ \"$VERCEL_ENV\" == \"production\" ]; then exit 1; else exit 0; fi"
  root_directory = "frontend/public/nuxt"
}

resource "vercel_project_domain" "skill_sheet" {
  project_id = vercel_project.skill_sheet.id
  domain     = "skill-sheet-blush.vercel.app"
}

resource "vercel_project_environment_variable" "skill_sheet" {
  project_id = vercel_project.skill_sheet.id
  target     = ["production", "preview", "development"]
  key        = "API_ENDPOINT"
  value      = "https://skill-sheet-production.up.railway.app/query"
}

resource "vercel_project" "skill_sheet_admin" {
  name           = "skill-sheet-admin"
  framework      = "nextjs"
  git_repository = local.git_repository
  ignore_command = "if [ \"$VERCEL_ENV\" == \"production\" ]; then exit 1; else exit 0; fi"
  root_directory = "frontend/admin/nextjs"
}

resource "vercel_project_domain" "skill_sheet_admin" {
  project_id = vercel_project.skill_sheet_admin.id
  domain     = "skill-sheet-admin.vercel.app"
}

locals {
  environment_variables = {
    NEXT_PUBLIC_AUTH0_DOMAIN       = var.auth0_domain
    NEXT_PUBLIC_AUTH0_CLIENT_ID    = var.auth0_client_id
    NEXT_PUBLIC_AUTH0_REDIRECT_URI = vercel_project_domain.skill_sheet_admin.domain
  }
}

resource "vercel_project_environment_variable" "skill_sheet_admin" {
  for_each = local.environment_variables

  project_id = vercel_project.skill_sheet_admin.id
  target     = ["production", "preview", "development"]
  key        = each.key
  value      = each.value
}
