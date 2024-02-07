resource "railway_service" "skill_sheet" {
  project_id     = railway_project.ogoshikazuki.id
  name           = "skill-sheet"
  config_path    = "infrastructure/railway/railway.json"
  root_directory = "/backend/go"
  source_repo    = "ogoshikazuki/skill-sheet"
}

resource "railway_deployment_trigger" "skill_sheet" {
  service_id     = railway_service.skill_sheet.id
  environment_id = railway_environment.production.id
  repository     = "ogoshikazuki/skill-sheet"
  branch         = "main"
}

resource "railway_service_domain" "skill_sheet" {
  service_id     = railway_service.skill_sheet.id
  environment_id = railway_environment.production.id
  subdomain      = "skill-sheet-production"
}

locals {
  variables = {
    CORS_ALLOWED_ORIGINS = "https://skill-sheet-blush.vercel.app"
    POSTGRES_HOST        = "postgres.railway.internal"
    POSTGRES_DBNAME      = "$${{Postgres.PGDATABASE}}"
    POSTGRES_USER        = "$${{Postgres.PGUSER}}"
    POSTGRES_PASSWORD    = "$${{Postgres.PGPASSWORD}}"
  }
}

resource "railway_variable" "skill_sheet" {
  for_each = local.variables

  service_id     = railway_service.skill_sheet.id
  environment_id = railway_environment.production.id
  name           = each.key
  value          = each.value
}
