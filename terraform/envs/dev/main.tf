provider "aws" {
  region = "us-east-1"
}

module "organizations_service" {
  source = "../../modules/organizations-service"
  prefix = "tfom-org-service"
}

module "modules_service" {
  source = "../../modules/modules-service"
  prefix = "tfom-modules-service"
}