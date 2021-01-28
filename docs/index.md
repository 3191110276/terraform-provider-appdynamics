---
page_title: "Provider: AppDynamics"
subcategory: ""
description: |-
  Terraform provider for interacting with AppDynamics API.
---

# AppDynamics Provider

The AppDynamics provider is used to interact with the AppDynamics APM. This provider only provides minimal capabilities to help when provisioning applications.

## Example Usage

Do not keep your authentication password in HCL for production environments, use Terraform environment variables.

```terraform
provider "appdynamics" {
  base_url = "https://example.saas.appdynamics.com"
  token = "your_token_goes_here"
}
```
