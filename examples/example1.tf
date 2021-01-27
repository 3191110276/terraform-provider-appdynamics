terraform {
  required_providers {
    appdynamics = {
      version = "0.0.4"
      source = "3191110276/appdynamics"
    }
  }
}

provider "appdynamics" {
  base_url = "dos"
  token = "test123"
}

resource "appdynamics_application" "brewery" {
  name = "brewery"
  description = "test"
}

output "new_app" {
  value = appdynamics_application.brewery
}
