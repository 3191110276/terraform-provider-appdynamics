---
page_title: "application Resource - terraform-provider-appdynamics"
subcategory: ""
description: |-
  The appdynamics_db_collector resource allows you to configure an AppDynamics application.
---

# Resource `appdynamics_db_collector`

## Example Usage

```terraform
resource "appdynamics_db_collector" "example" {
  name = "DB NAME"
  type = "MYSQL"
  hostname = "hostname"
  port = "80"
  username = "root"
  password = "root"
  agent_name = "agent"
}
```

## Argument Reference

- `name` - (Required) Name of the application in AppDynamics

- `type` - (Required) Type of the database

- `hostname` - (Required) Hostname of the database

- `port` - (Required) Port of the database

- `username` - (Required) Username of the database

- `password` - (Required) Password of the database

- `agent_name` - (Required) Agent name used for connecting to the database
