resource "vercel_project" "skill_sheet" {
  name      = "skill-sheet"
  framework = "nuxtjs"
  git_repository = {
    type              = "github"
    repo              = "ogoshikazuki/skill-sheet"
    production_branch = "main"
  }
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
