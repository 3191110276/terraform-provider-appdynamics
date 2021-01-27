terraform {
  required_providers {
    appdynamics = {
      #versions = ["0.3"]
      source = "github.com/3191110276/terraform-provider-appdynamics"
    }
  }
}

provider "appdynamics" {
  username = "dos"
  password = "test123"
}

resource "appd_app" "brewery" {
  items {
    coffee {
      id = 3
    }
    quantity = 2
  }
  items {
    coffee {
      id = 2
    }
    quantity = 2
  }
}

output "new_app" {
  value = appd_app.brewery
}
