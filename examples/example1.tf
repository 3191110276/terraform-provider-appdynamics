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
