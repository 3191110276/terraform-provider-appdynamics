---
page_title: "application Resource - terraform-provider-appdynamics"
subcategory: ""
description: |-
  The order resource allows you to configure an AppDynamics application.
---

# Resource `appdynamics_apm_application`

## Example Usage

```terraform
resource "appdynamics_apm_application" "example" {
  name = "appname"
  description = "appdescription"
}
```

## Argument Reference

- `name` - (Required) Name of the application in AppDynamics

- `description` - (Optional) Description of the application in AppDynamics
