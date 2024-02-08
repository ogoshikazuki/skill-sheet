resource "railway_project" "ogoshikazuki" {
  name = "ogoshikazuki"
}

resource "railway_environment" "production" {
  project_id = railway_project.ogoshikazuki.id
  name       = "production"
}
