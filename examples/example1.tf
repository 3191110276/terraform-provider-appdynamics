terraform {
  required_providers {
    appdynamics = {
      #versions = ["0.3"]
      source = "github.com/3191110276/appdynamics"
    }
  }
}

provider "appdynamics" {
  base_url = "dos"
  token = "test123"
}

resource "appd_app" "brewery" {
  name = "brewery"
  description = "test"
}

output "new_app" {
  value = appd_app.brewery
}
