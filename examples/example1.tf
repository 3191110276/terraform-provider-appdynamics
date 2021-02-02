terraform {
  required_providers {
    appdynamics = {
      version = "0.0.66"
      source = "3191110276/appdynamics"
    }
  }
}

provider "appdynamics" {
  base_url = "https://ceer.saas.appdynamics.com"
  token = "eyJraWQiOiJhNTEwOWE5ZC04NWRkLTRmZWItOTE4NS00ZGE1NzZjMjExZDciLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJBcHBEeW5hbWljcyIsImF1ZCI6IkFwcERfQVBJcyIsImp0aSI6IlZEV1B5eXNXU1p0cGhqejR0NXduSkEiLCJzdWIiOiJtaW1hdXJlcmFwaSIsImlkVHlwZSI6IkFQSV9DTElFTlQiLCJpZCI6IjlhNjhkNzJhLWNkMjQtNDM3MS04YmFmLWI1YzhlYTcwZmRjZSIsImFjY3RJZCI6ImE1MTA5YTlkLTg1ZGQtNGZlYi05MTg1LTRkYTU3NmMyMTFkNyIsInRudElkIjoiYTUxMDlhOWQtODVkZC00ZmViLTkxODUtNGRhNTc2YzIxMWQ3IiwiYWNjdE5hbWUiOiJjZWVyIiwidGVuYW50TmFtZSI6IiIsImZtbVRudElkIjpudWxsLCJhY2N0UGVybSI6W10sInJvbGVJZHMiOltdLCJpYXQiOjE2MTE1MDQ4NzYsIm5iZiI6MTYxMTUwNDc1NiwiZXhwIjoxNjQzMDQwODc2LCJ0b2tlblR5cGUiOiJBQ0NFU1MifQ.TMqYCqBgL1RJGZ1KFyLIDXc-KY3w7YcxQDauSgmcCT8"
  username = "mimaurer@ceer"
  password = "mimaurer"
}

resource "appdynamics_apm_application" "brewery" {
  name = "brewery"
  description = "Demo Application"
}

resource "appdynamics_eum_application" "brewery" {
  name = "brewery_eum"
  description = "test"
}

resource "appdynamics_db_collector" "brewery" {
  name = "Inventory DB"
  type = "MYSQL"
  hostname = "brewery-inventorydb-service.app"
  port = "80"
  username = "root"
  password = "root"
  agent_name = "brewery"
}

output "new_db" {
  value = appdynamics_db_collector.brewery
}
