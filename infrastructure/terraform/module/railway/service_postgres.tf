resource "railway_service" "postgres" {
  project_id   = railway_project.ogoshikazuki.id
  name         = "Postgres"
  source_image = "ghcr.io/railwayapp-templates/postgres-ssl:latest"
  volume = {
    mount_path = "/var/lib/postgresql/data"
    name       = "pgdata"
  }
}

resource "railway_tcp_proxy" "postgres" {
  service_id       = railway_service.postgres.id
  environment_id   = railway_environment.production.id
  application_port = 5432
}
